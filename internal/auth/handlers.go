package auth

import (
	"encoding/json"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/captcha"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type LoginType struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hcaptcha string `json:"h-captcha-response"`
	Gcaptcha string `json:"g-recaptcha-response"`
}

type RegisterType struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hcaptcha string `json:"h-captcha-response"`
	Gcaptcha string `json:"g-recaptcha-response"`
}

func Register(c fiber.Ctx) error {
	fmt.Println("User sent register request.")
	var user RegisterType
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	fmt.Println(user)

	if validName := validUsername(user.Username); !validName {
		return c.JSON(fiber.Map{"error": "This username is not valid"})
	}

	if validPass := validPassword(user.Password); !validPass {
		return c.JSON(fiber.Map{"error": "Your password must be 8-32 characters and can only contain letters and numbers."})
	}

	fmt.Printf("Signup: %s:%s\n", user.Username, user.Password)
	// require captcha before checking password hash.

	success, captchaErr := captcha.VerifyHCaptcha(user.Hcaptcha)
	if captchaErr != nil {
		log.Fatalf("Error verifying hCaptcha: %v", captchaErr)
	}

	if !success {
		return util.CaptchaFailed(c)
	}

	fmt.Println("hCaptcha verification succeeded")

	exists, err := usernameExists(user.Username)
	if err {
		fmt.Println("user exists internal")
		return util.InternalServerError(c)
	}
	if exists {
		return c.JSON(fiber.Map{"error": "User already exists"})
	}

	hashedPass, passErr := hashPassword(user.Password)
	if passErr != nil {
		fmt.Println("hashPassword internal")
		return util.InternalServerError(c)
	}

	userRegistered := RegisterUserDatabase(user.Username, hashedPass)

	if !userRegistered {
		fmt.Println("user not registered internal")
		return util.InternalServerError(c)
	}

	sessionErr := createSession(user.Username, c)
	if sessionErr != nil {
		// if the session fails to create on register then redirect to auth
		// frontend must redirect to auth
		return c.JSON(fiber.Map{"message": "User registered, Login required."})
	} else {
		// redirect to profile page
		return c.JSON(fiber.Map{"message": "User registered"})
	}

}

func Login(c fiber.Ctx) error {
	fmt.Println("User sent auth request.")
	var user LoginType
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	fmt.Println(user)
	if user.Username == "" || !validUsername(user.Username) {
		return util.InvalidCredentials(c)
	}

	if !validPassword(user.Password) {
		return util.InvalidCredentials(c)
	}

	storedPassword, valid := LoginGetPassword(user.Username)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to verify user, are you sure your username is correct?"})
	}

	// require captcha before checking password hash.

	success, captchaErr := captcha.VerifyHCaptcha(user.Hcaptcha)
	if captchaErr != nil {
		log.Fatalf("Error verifying hCaptcha: %v", captchaErr)
	}

	if !success {
		return util.CaptchaFailed(c)
	}

	fmt.Println("hCaptcha verification succeeded")

	match := verifyPassword(storedPassword, user.Password)

	if !match {
		return util.InvalidCredentials(c)
	}

	sessionErr := createSession(user.Username, c)
	if sessionErr == nil {
		return c.JSON(fiber.Map{"message": "Logged in"})
	}

	// this should not happen, maybe memory full?

	return util.InternalServerError(c)

}

func Logout(c fiber.Ctx) error {
	fmt.Println("User sent logout request.")
	sess, _ := UserSessionStore.Get(c)
	sess.Destroy()
	return c.JSON(fiber.Map{"message": "Logged out"})
}

package auth

import (
	"encoding/json"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/api/auth/captcha"
	"github.com/connect-web/Low-Latency-API/internal/db/auth"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

func Logout(c fiber.Ctx) error {
	fmt.Println("User sent logout request.")
	sess, _ := UserSessionStore.Get(c)
	sess.Destroy()
	return c.JSON(fiber.Map{"message": "Logged out"})
}

func Register(c fiber.Ctx) error {
	var user model.RegisterType
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if validName := auth.ValidUsername(user.Username); !validName {
		return c.JSON(fiber.Map{"error": "This username is not valid."})
	}

	if validPass := auth.ValidPassword(user.Password); !validPass {
		return c.JSON(fiber.Map{"error": "Your password must be 8-32 characters and can only contain letters and numbers."})
	}
	// require captcha before checking password hash.

	if user.Hcaptcha == "" {
		return util.CaptchaFailed(c)
	}

	success, captchaErr := captcha.VerifyHCaptcha(user.Hcaptcha)
	if captchaErr != nil {
		fmt.Printf("Error verifying hCaptcha: %v", captchaErr)
		return util.InternalServerError(c)
	}

	if !success {
		return util.CaptchaFailed(c)
	}

	exists, err := auth.UsernameExists(user.Username)
	if err {
		fmt.Println("user exists internal")
		return util.InternalServerError(c)
	}
	if exists {
		return c.JSON(fiber.Map{"error": "Username already exists."})
	}

	hashedPass, passErr := auth.HashPassword(user.Password)
	if passErr != nil {
		fmt.Println("hashPassword internal")
		return util.InternalServerError(c)
	}

	userRegistered := auth.RegisterUserDatabase(user.Username, hashedPass)

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
	var user model.LoginType
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if user.Username == "" || !auth.ValidUsername(user.Username) {
		return util.InvalidCredentials(c)
	}

	if !auth.ValidPassword(user.Password) {
		return util.InvalidCredentials(c)
	}

	// require captcha before checking password hash.

	if user.Hcaptcha == "" {
		return util.CaptchaFailed(c)
	}

	success, captchaErr := captcha.VerifyHCaptcha(user.Hcaptcha)
	if captchaErr != nil {
		fmt.Printf("Error verifying hCaptcha: %v", captchaErr)
		return util.InternalServerError(c)
	}

	if !success {
		return util.CaptchaFailed(c)
	}

	storedPassword, valid := auth.LoginGetPassword(user.Username)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to verify user, are you sure your username is correct?"})
	}
	fmt.Println("Fetched password")

	match := auth.VerifyPassword(storedPassword, user.Password)

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

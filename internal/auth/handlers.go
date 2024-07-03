package auth

import (
	"encoding/json"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

type LoginType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterType struct {
	Username string `json:"username"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

func Register(c fiber.Ctx) error {
	fmt.Println("User sent register request.")
	var user RegisterType
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if validName := validUsername(user.Username); !validName {
		return c.JSON(fiber.Map{"error": "This username is not valid"})
	}

	if validPass := validPassword(user.Password); !validPass {
		return c.JSON(fiber.Map{"error": "Your password must be 8-32 characters and can only contain letters and numbers."})
	}

	fmt.Printf("Signup: %s:%s\n", user.Username, user.Password)
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

	return c.JSON(fiber.Map{"message": "User registered"})
}

func Login(c fiber.Ctx) error {
	fmt.Println("User sent login request.")
	var user LoginType
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if user.Username == "" || !validUsername(user.Username) {
		return util.InvalidCredentials(c)
	}

	if !validPassword(user.Password) {
		return util.InvalidCredentials(c)
	}

	storedPassword, valid := LoginGetPassword(user.Username)
	if !valid {
		return util.InternalServerError(c)
	}

	match := verifyPassword(storedPassword, user.Password)

	if !match {
		return util.InvalidCredentials(c)
	}

	// Use user session store for authenticated users
	/*
		sess, err := userSessionStore.Get(c)
		if err != nil {
			return err
		}
		sess.Set("username", user.Username)
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session save error"})
		}

	*/

	return c.JSON(fiber.Map{"message": "Logged in"})
}

func Logout(c fiber.Ctx) error {
	fmt.Println("User sent logout request.")
	sess, _ := userSessionStore.Get(c)
	sess.Destroy()
	return c.JSON(fiber.Map{"message": "Logged out"})
}

func Protected(c fiber.Ctx) error {
	fmt.Println("User accessed protected route.")
	sess, err := userSessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session retrieval error"})
	}
	if sess.Get("username") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	return c.Next()
}
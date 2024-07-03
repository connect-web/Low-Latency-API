package auth

import "github.com/gofiber/fiber/v3"

func createSession(username string, c fiber.Ctx) error {
	sess, err := UserSessionStore.Get(c)
	if err != nil {
		return err
	}
	sess.Set("username", username)
	return sess.Save()
}

package handlers

import "github.com/gofiber/fiber/v2"

// login godoc
//
//	@Summary		Login in a user and create a session
//	@Description	Once "logged in" a users auth token will last for one hour
//	@Tags			users
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/api/v1/login [get]
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

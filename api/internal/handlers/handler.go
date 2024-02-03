package handlers

import "github.com/gofiber/fiber/v2"

// Placeholder file

// Mock handler

// handler godoc
//
//	@Summary		Mock handler
//	@Description	Mock handler
//	@Tags			handlers
//	@Produce		json
//	@Success		200
//	@Failure		500
//	@Router			/api/v1/handler [get]
func Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

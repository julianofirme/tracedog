package ingestion

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julianofirme/tracedog/internal/core"
	"github.com/julianofirme/tracedog/internal/processor"
	"github.com/rs/zerolog/log"
)

func RegisterRoutes(app *fiber.App) {
	app.Post("/events", handleEvent)
}

func handleEvent(c *fiber.Ctx) error {
	var payload core.EventPayload

	if err := c.BodyParser(&payload); err != nil {
		log.Error().Err(err).Msg("invalid json")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON",
		})
	}

	if err := ValidateEvent(&payload); err != nil {
		log.Warn().Err(err).Msg("validation failed")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	processor.EventQueue(payload)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "event received",
	})
}

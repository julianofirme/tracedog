package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/julianofirme/tracedog/internal/config"
	"github.com/julianofirme/tracedog/internal/ingestion"
	"github.com/julianofirme/tracedog/internal/processor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := config.LoadConfig()
	log.Info().Msg("Starting TraceDog...")

	app := fiber.New()

	ingestion.RegisterRoutes(app)
	processor.InitQueue(1000)
	processor.StartWorker()

	ingestion.RegisterRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("TraceDog is alive.")
	})

	addr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Info().Msgf("Listening on port %s ", cfg.AppPort)
	if err := app.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

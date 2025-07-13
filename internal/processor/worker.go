package processor

import (
	"github.com/rs/zerolog/log"
)

func StartWorker() {
	go func() {
		for event := range GetQueue() {
			log.Info().
				Str("event_name", event.EventName).
				Str("user_id", event.UserID).
				Interface("props", event.Props).
				Msg("event processed by worker")
		}
	}()
}

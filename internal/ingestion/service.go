package ingestion

import (
	"errors"

	"github.com/julianofirme/tracedog/internal/core"
)

func ValidateEvent(e *core.EventPayload) error {
	if e.EventName == "" {
		return errors.New("event_name is required")
	}

	if e.UserID == "" {
		return errors.New("user_id is required")
	}

	if e.Timestamp.IsZero() {
		return errors.New("timestamp is required")
	}

	return nil
}

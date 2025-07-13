package core

import "time"

type EventPayload struct {
	EventName string                 `json:"event_name"`
	Timestamp time.Time              `json:"timestamp"`
	UserID    string                 `json:"user_id"`
	Props     map[string]interface{} `json:"props"`
}

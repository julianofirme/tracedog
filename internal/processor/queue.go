package processor

import (
	"github.com/julianofirme/tracedog/internal/core"
)

var eventQueue chan core.EventPayload

func InitQueue(size int) {
	eventQueue = make(chan core.EventPayload, size)
}

func EventQueue(e core.EventPayload) {
	eventQueue <- e
}

func GetQueue() <-chan core.EventPayload {
	return eventQueue
}

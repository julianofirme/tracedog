package processor

import (
	"github.com/julianofirme/tracedog/internal/store"
)

func StartWorker() {
	go func() {
		for event := range GetQueue() {
			store.GetStore().SaveEvent(event)
		}
	}()
}

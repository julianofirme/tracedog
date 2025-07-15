package store

import (
	"sync"

	"github.com/julianofirme/tracedog/internal/core"
)

type MemoryStore struct {
	mu       sync.RWMutex
	events   []core.EventPayload
	counts   map[string]int
	usersSet map[string]struct{}
}

var instance *MemoryStore
var once sync.Once

func GetStore() *MemoryStore {
	once.Do(func() {
		instance = &MemoryStore{
			events:   make([]core.EventPayload, 0),
			counts:   make(map[string]int),
			usersSet: make(map[string]struct{}),
		}
	})
	return instance
}

func (s *MemoryStore) SaveEvent(e core.EventPayload) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.events = append(s.events, e)
	s.counts[e.EventName]++
	s.usersSet[e.UserID] = struct{}{}
}

func (s *MemoryStore) CountEvents() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.events)
}

func (s *MemoryStore) EventCounts() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	copy := make(map[string]int)
	for k, v := range s.counts {
		copy[k] = v
	}
	return copy
}

func (s *MemoryStore) UniqueUserCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.usersSet)
}

package observer

import (
	"sync"

	"github.com/hieua1/logger"
)

type Subject interface {
	RegisterObserver(obs Observer)
	UnregisterObserver(obs Observer)
	NotifyAll(data interface{})
}

type BaseSubject struct {
	observers []Observer
	mu        sync.Mutex
}

func (s *BaseSubject) RegisterObserver(obs Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, obs)
}

func (s *BaseSubject) UnregisterObserver(obs Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for index, o := range s.observers {
		if o == obs {
			s.observers = append(s.observers[:index], s.observers[index+1:]...)
			return
		}
	}
	logger.S().Warn("Unregistered observer failed: Couldn't find observer in observers list.")
}

func (s *BaseSubject) NotifyAll(data interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, obs := range s.observers {
		obs.OnNotify(data)
	}
}
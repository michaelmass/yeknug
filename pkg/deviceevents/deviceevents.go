package deviceevents

import (
	"github.com/go-vgo/robotgo"
)

const (
	KeyDown   = 3
	MouseDown = 8
)

type Event struct {
	Kind int
}

type KeyEvents struct {
	eventsCh chan Event
}

func New() (*KeyEvents, error) {
	eventsCh := make(chan Event, 10)

	go listen(eventsCh)

	return &KeyEvents{
		eventsCh: eventsCh,
	}, nil
}

func (keyEvents *KeyEvents) Events() <-chan Event {
	return keyEvents.eventsCh
}

func listen(eventsCh chan<- Event) {
	s := robotgo.EventStart()

	for e := range s {
		eventsCh <- Event{
			Kind: int(e.Kind),
		}
	}
}

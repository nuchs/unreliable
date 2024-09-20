package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

type Top struct {
	state      State
	volitility int
	period     time.Duration
	Down       chan State
	Up         chan any
}

func NewTop(volitility int, period time.Duration) *Top {
	return &Top{
		newState(0, 0),
		volitility,
		period,
		make(chan State, 10),
		make(chan any, 10),
	}
}

func (t *Top) Run(ctx context.Context) {
	log.Printf(" | TOPP | Started\n")

	ticker := time.NewTicker(t.period * time.Microsecond)

	for {
		select {
		case <-ctx.Done():
			log.Printf(" | TOPP | Terminated: %+v\n", t.state)
			return
		case <-t.Up:
			t.Down <- t.state
		case <-ticker.C:
			if t.volitility < rand.Intn(100) {
				old := t.state
				t.state = t.state.next()
				log.Printf(" | TOPP | State %+v -> %+v\n", old, t.state)
			}
		}
	}
}

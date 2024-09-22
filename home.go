package main

import (
	"context"
	"log"
	"time"
)

type Home struct {
	state        State
	down         *Middle
	up           *Middle
	top          *Top
	topPeriod    time.Duration
	bottomPeriod time.Duration
}

var blank = State{0, 0, 0}

func NewHome(down, up *Middle, top *Top, topPeriod, bottomPeriod time.Duration) *Home {
	return &Home{blank, down, up, top, topPeriod, bottomPeriod}
}

func (h *Home) Run(ctx context.Context) {
	log.Printf(" | HOME | Started\n")
	topPoll := time.NewTicker(time.Microsecond * h.topPeriod)
	bottomPoll := time.NewTicker(time.Microsecond * h.bottomPeriod)
	clock := 0

	for {
		select {
		case <-ctx.Done():
			log.Printf(" | HOME | Terminated\n")
			return
		case <-topPoll.C:
			h.top.Up <- 0
		case s := <-h.top.Down:
			if s.Value != h.state.Value {
				log.Printf(" | HOME | State %+v -> %+v\n", h.state, s)
				h.state = s
			}
		case <-bottomPoll.C:
			clock++
			msg := Msg{query, 0, blank, clock}
			h.down.Send <- msg
			log.Printf(" | HOME | Sent request: %+v\n", msg)
		case msg := <-h.up.Recv:
			log.Printf(" | HOME | Received response: %+v\n", msg)
			if msg.State.Value != h.state.Value {
				clock++
				msgOut := Msg{update, msg.Hash, h.state, clock}
				h.down.Send <- msgOut
				log.Printf(" | HOME | Sent request: %+v\n", msgOut)
			}
		}
	}
}

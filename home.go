package main

import (
	"context"
	"log"
	"time"
)

type Home struct {
	state  State
	down   *Middle
	up     *Middle
	top    *Top
	period time.Duration
}

func NewHome(down, up *Middle, top *Top, period time.Duration) *Home {
	return &Home{blank, down, up, top, period}
}

func (h *Home) Run(ctx context.Context) {
	log.Printf(" | HOME | Started\n")
	poll := time.NewTicker(time.Microsecond * h.period)

	for {
		select {
		case <-ctx.Done():
			log.Printf(" | HOME | Terminated\n")
			return
		case <-poll.C:
			h.top.Up <- 0
		case s := <-h.top.Down:
			if s.Value != h.state.Value {
				log.Printf(" | HOME | State %+v -> %+v\n", h.state, s)
				h.state = s
			}
		case msg := <-h.up.Recv:
			log.Printf(" | HOME | Received request: %+v\n", msg)
			msgOut := Msg{update, h.state, msg.Session, msg.Clock}
			h.down.Send <- msgOut
			log.Printf(" | HOME | Sent response: %+v\n", msgOut)
		}
	}
}

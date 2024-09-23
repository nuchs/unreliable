package main

import (
	"context"
	"log"
)

type Bottom struct {
	state  State
	down   chan Msg
	up     chan Resp
	target int
}

type Resp struct {
	State   State
	Session int
	Clock   int
}

func NewBottom(target int) *Bottom {
	return &Bottom{
		newState(0, 0),
		make(chan Msg, 10),
		make(chan Resp, 10),
		target,
	}
}

func (b *Bottom) Run(ctx context.Context, cancel context.CancelFunc) {
	log.Printf(" | BOTT | Started\n")

	for {
		select {
		case <-ctx.Done():
			log.Printf(" | BOTT | Terminated: %+v\n", b.state)
			return
		case msg := <-b.down:
			b.Handle(msg, cancel)
		}
	}
}

func (b *Bottom) Handle(msg Msg, cancel context.CancelFunc) {
	switch msg.Kind {
	case status:
		b.up <- Resp{b.state, msg.Session, msg.Clock}
	case update:
		log.Printf(" | BOTT | Updating %+v -> %+v\n", b.state, msg.State)
		if b.state.Version >= msg.State.Version {
			log.Printf(" | BOTT | Error detected version must increase on update\n")
		}
		b.state = msg.State
		if b.state.Value >= b.target {
			log.Printf(" | BOTT | Target reached %+v\n", b.state)
			cancel()
		}
	default:
		log.Printf(" | BOTT | Error unknown message %+v\n", msg)
		cancel()
	}
}

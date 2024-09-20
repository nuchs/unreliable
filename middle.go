package main

import (
	"context"
	"log"
	"math/rand"
)

type Middle struct {
	name  string
	Send  chan Msg
	Recv  chan Msg
	delay chan Msg
}

func NewMiddle(name string) *Middle {
	return &Middle{
		name,
		make(chan Msg, 1000),
		make(chan Msg, 1000),
		make(chan Msg, 1000),
	}
}

func (m *Middle) Run(ctx context.Context) {
	log.Printf(" | %s | Started\n", m.name)

	for {
		select {
		case <-ctx.Done():
			log.Printf(" | %s | Terminated\n", m.name)
			return
		case msg := <-m.Send:
			m.Enqueue(msg)
		}
	}
}

func (m *Middle) Enqueue(msg Msg) {
	switch rand.Intn(20) {
	case 0:
		log.Printf(" | %s | Discarded: %+v\n", m.name, msg)
	case 1:
		log.Printf(" | %s | Duplicated: %+v\n", m.name, msg)
		m.Recv <- msg
		m.Recv <- msg
	case 2:
		log.Printf(" | %s | Delayed: %+v\n", m.name, msg)
		m.delay <- msg
	case 3:
		log.Printf(" | %s | Deliver %d delayed messages\n", m.name, len(m.delay))
		for {
			select {
			case delayed := <-m.delay:
				m.Recv <- delayed
			default:
				m.Recv <- msg
				return
			}
		}
	default:
		m.Recv <- msg
	}
}

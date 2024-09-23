package main

import (
	"context"
	"log"
	"time"
)

type Away struct {
	clock   int
	session int
	down    *Middle
	up      *Middle
	bottom  *Bottom
	period  time.Duration
}

func NewAway(down, up *Middle, bottom *Bottom, period time.Duration) *Away {
	return &Away{0, 1, down, up, bottom, period}
}

func (a *Away) Run(ctx context.Context) {
	log.Printf(" | AWAY | Started\n")
	poll := time.NewTicker(time.Microsecond * a.period)

	for {
		select {
		case <-ctx.Done():
			log.Printf(" | AWAY | Terminated\n")
			return
		case <-poll.C:
			a.bottom.down <- Msg{status, blank, a.session, a.clock}
		case msg := <-a.bottom.up:
			a.HandleBottom(msg)
		case msg := <-a.down.Recv:
			a.HandleHome(msg)
		}
	}
}

func (a *Away) HandleHome(msg Msg) {
	log.Printf(" | AWAY | Received update: %+v\n", msg)
	if msg.Session != a.session {
		log.Printf(" | AWAY | Bad session: a(%+v) vs m(%+v)\n", a.session, msg.Session)
		return
	}

	if msg.Clock < a.clock {
		log.Printf(" | AWAY | Bad clock: a(%+v) vs m(%+v)\n", a.clock, msg.Clock)
		return
	}
	log.Printf(" | AWAY | Starting update: %+v\n", msg)
	a.clock++
	a.bottom.down <- msg
	a.bottom.down <- Msg{status, blank, a.session, a.clock}
	for {
		resp := <-a.bottom.up
		if resp.Clock == a.clock {
			log.Printf(" | AWAY | Update complete: %+v\n", resp)
			break
		}
	}
}

func (a *Away) HandleBottom(resp Resp) {
	if resp.Session != a.session {
		log.Printf(" | AWAY | Session mismatch: a(%+v) vs b(%+v)\n", a.session, resp.Session)
		return
	}
	msg := Msg{status, resp.State, resp.Session, resp.Clock}
	a.up.Send <- msg
	log.Printf(" | AWAY | Sent status: %+v\n", msg)
}

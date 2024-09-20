package main

import (
	"context"
	"log"
)

type Away struct {
	hash   int
	down   *Middle
	up     *Middle
	bottom *Bottom
}

func NewAway(down, up *Middle, bottom *Bottom) *Away {
	return &Away{-1, down, up, bottom}
}

func (a *Away) Run(ctx context.Context) {
	log.Printf(" | AWAY | Started\n")

	for {
		select {
		case <-ctx.Done():
			log.Printf(" | AWAY | Terminated\n")
			return
		case msg := <-a.down.Recv:
			a.HandleRequest(msg)
		case msg := <-a.bottom.up:
			a.HandleResponse(msg)
		}
	}
}

func (a *Away) HandleRequest(msg Msg) {
	log.Printf(" | AWAY | Received request: %+v\n", msg)

	if msg.Kind == update {
		if msg.Hash != a.hash {
			log.Printf(" | AWAY | Bad hash msg:%+v vs away:%+v\n", msg.Hash, a.hash)
			return
		}
		a.bottom.down <- msg
		a.bottom.down <- Msg{query, 0, msg.State, msg.Clock}
		resp := <-a.bottom.up
		a.HandleResponse(resp)
	} else {
		a.bottom.down <- msg
	}
}

func (a *Away) HandleResponse(resp Resp) {
	if a.hash != resp.State.Hash {
		a.hash = resp.State.Hash
		log.Printf(" | AWAY | New hash: %+v\n", a.hash)
	}

	msg := Msg{status, a.hash, resp.State, resp.Clock}
	a.up.Send <- msg
	log.Printf(" | AWAY | Sent response: %+v\n", msg)
}

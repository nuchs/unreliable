package main

import (
	"context"
	"log"
	"sync"
	"time"
)

type Config struct {
	volitility int
	topPeriod  time.Duration
	topPoll    time.Duration
	bottomPoll time.Duration
	target     int
}

func main() {
	config := Config{
		volitility: 1,
		topPeriod:  3 * time.Microsecond,
		topPoll:    5 * time.Microsecond,
		bottomPoll: 1 * time.Microsecond,
		target:     1000000,
	}

	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Printf(" | MAIN | Hello! It's a me, Nodap!\n")

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	top := NewTop(config.volitility, config.topPeriod)
	down := NewMiddle("DOWN")
	up := NewMiddle(" UP ")
	bottom := NewBottom(config.target)
	home := NewHome(down, up, top, config.topPoll, config.bottomPoll)
	away := NewAway(down, up, bottom)

	log.Printf(" | MAIN | Start simulation\n")
	start(func() { bottom.Run(ctx, cancel) }, &wg)
	start(func() { away.Run(ctx) }, &wg)
	start(func() { up.Run(ctx) }, &wg)
	start(func() { down.Run(ctx) }, &wg)
	start(func() { top.Run(ctx) }, &wg)
	start(func() { home.Run(ctx) }, &wg)

	log.Printf(" | MAIN | Waiting for termination\n")
	wg.Wait()
	log.Printf(" | MAIN | Ok lady, I love you buhbye\n")
}

func start(run func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		run()
	}()
}

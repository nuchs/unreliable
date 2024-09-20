package main

import "math/rand"

type Top struct {
	state      int
	volitility int
	log        func(string, ...any)
}

func NewTop(volitility int, log func(string, ...any)) *Top {
	return &Top{0, volitility, log}
}

func (t *Top) Tick() {
	if t.volitility < rand.Intn(100) {
		old := t.state
		t.state += delta()
		t.log("State %d -> %d", old, t.state)
	}
}

func delta() int {
	base := rand.Intn(5) - 2
	if base == 0 {
		base += 3
	}

	return base
}

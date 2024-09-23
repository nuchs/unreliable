package main

import "math/rand"

type State struct {
	Value   int
	Version int
}

var blank = State{0, 0}

func newState(value int, version int) State {
	return State{value, version}
}

func (s State) next() State {
	return newState(s.Value+delta(), s.Version+1)
}

func delta() int {
	base := rand.Intn(5) - 2
	if base == 0 {
		base += 3
	}

	return base
}

type Kind string

const (
	status Kind = "status"
	update Kind = "update"
)

type Msg struct {
	Kind    Kind
	State   State
	Session int
	Clock   int
}

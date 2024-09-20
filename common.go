package main

import "math/rand"

type State struct {
	Value   int
	Version int
	Hash    int
}

func newState(value int, version int) State {
	return State{value, version, (value * 1000) + 111}
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
	query  Kind = "query"
	status Kind = "status"
	update Kind = "update"
)

type Msg struct {
	Kind  Kind
	Hash  int
	State State
	Clock int
}

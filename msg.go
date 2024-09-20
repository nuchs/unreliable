package main

type Kind int

const (
	query Kind = iota
	status
	update
)

type Msg struct {
	Kind  Kind
	Hash  int
	State int
}

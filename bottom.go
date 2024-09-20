package main

type Bottom struct {
	state int
	log   func(string, ...any)
}

func NewBottom(log func(string, ...any)) *Bottom {
	return &Bottom{0, log}
}

func (b *Bottom) Query() int {
	return b.state
}

func (b *Bottom) Update(state int) {
	old := b.state
	b.state = state
	b.log("State %d -> %d", old, state)
}

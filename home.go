package main

type Home struct {
	state   int
	send    func(Msg)
	receive func() Msg
	log     func(string, ...any)
}

func NewHome(send func(Msg), receive func() Msg, log func(string, ...any)) *Home {
	return &Home{0, send, receive, log}
}

func (h *Home) Tick() {
}

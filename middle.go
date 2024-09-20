package main

type Middle struct {
	log func(string, ...any)
}

func NewMiddle(log func(string, ...any)) *Middle {
	return &Middle{log}
}

func (m *Middle) send(msg Msg) {

}

func (m *Middle) receive(msg Msg) {

}

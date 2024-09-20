package main

import "fmt"

type LogManager struct {
	time int
}

func (lm *LogManager) NewLogger(name string) func(string, ...any) {
	return func(msg string, args ...any) {
		formatted := fmt.Sprintf(msg, args...)
		fmt.Printf("%6d | %4s | %s", lm.time, name, formatted)
	}
}

func (lm *LogManager) Tick() {
	lm.time++
}

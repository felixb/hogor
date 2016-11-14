package main

import (
	"time"
)

const (
	max_wait = time.Second * 5
)

type WaitingState struct{}

func NewWaitingState() *WaitingState {
	s := WaitingState{}
	return &s
}

func (s *WaitingState) Enter() State {
	return nil
}

func (s *WaitingState) Event(pin uint, value uint) State {
	if pin == gpioSwitch && value == gpioSwitchOff {
		return NewOffState()
	} else {
		return nil
	}
}

func (s *WaitingState) String() string {
	return "Waiting"
}

func (s *WaitingState) Tick(d time.Duration) State {
	// TODO blink status led
	if d < max_wait {
		return nil
	} else {
		return NewOnState()
	}
}

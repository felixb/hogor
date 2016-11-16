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

func (s *WaitingState) Enter(m StateMachine) {}

func (s *WaitingState) Event(m StateMachine, pin uint, value uint) {
	if pin == GPIO_SWITCH_PIN && value == GPIO_SWITCH_OFF {
		m.Transit(STATE_OFF)
	}
}

func (s *WaitingState) Leave(m StateMachine) {}

func (s *WaitingState) String() string {
	return "Waiting"
}


func (s *WaitingState) Tick(d time.Duration) State {
	// FIXME get your own timer
	// TODO blink status led
	if d < max_wait {
		return nil
	} else {
		return NewOnState()
	}
}

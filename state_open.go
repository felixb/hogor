package main

import (
	"time"
)

type OpenState struct{}

func NewOpenState() *OpenState {
	s := OpenState{}
	return &s
}
func (s *OpenState) Enter(m StateMachine)  {
	// TODO ring the bell
}

func (s *OpenState) Event(m StateMachine, pin uint, value uint)  {
	if pin == GPIO_SWITCH_PIN && value == GPIO_SWITCH_OFF {
		m.Transit(STATE_OFF)
	}
}

func (s *OpenState) Leave(m StateMachine) {}

func (s *OpenState) String() string {
	return "Open"
}

func (s *OpenState) Tick(d time.Duration) State {
	// FIXME get your own timer
	if d < max_wait {
		return nil
	} else {
		return NewOnState()
	}
}

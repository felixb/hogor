package main

import (
	"log"
)

// The initial state does nothing
// but transit to the next state depending of inital switch values
type InititalState struct {
	switchValue uint
	gateValue   uint
}

func NewInititalState(switchValue, gateValue uint) State {
	s := InititalState{switchValue, gateValue}
	return &s
}

func (s *InititalState) Enter(m StateMachine) {
	log.Printf("Initial switch value: %d", s.switchValue)
	log.Printf("Initial gate value: %d", s.gateValue)

	if s.switchValue == GPIO_SWITCH_OFF {
		m.Transit(STATE_OFF)
	} else {
		m.Transit(STATE_WAITING)
	}
}

func (s *InititalState) Event(m StateMachine, pin uint, value uint) {}

func (s *InititalState) Leave(m StateMachine) {}

func (s *InititalState) String() string {
	return "Initital"
}

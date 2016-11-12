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

func (s *InititalState) Enter(m *Machine) {
	log.Printf("Initial switch value: %d", s.switchValue)
	log.Printf("Initial gate value: %d", s.gateValue)

	if s.switchValue == gpioSwitchOff {
		m.Transit(NewOffState())
	} else {
		m.Transit(NewWaitingState())
	}
}

func (s *InititalState) Leave(m *Machine) {}

func (s *InititalState) Event(m *Machine, pin uint, value uint) {}

func (s *InititalState) String() string {
	return "InititalState"
}

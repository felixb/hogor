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

func (s *InititalState) Enter() State {
	log.Printf("Initial switch value: %d", s.switchValue)
	log.Printf("Initial gate value: %d", s.gateValue)

	if s.switchValue == gpioSwitchOff {
		return NewOffState()
	} else {
		return NewWaitingState()
	}
}

func (s *InititalState) Event(pin uint, value uint) State {
	return nil
}

func (s *InititalState) String() string {
	return "Initital"
}

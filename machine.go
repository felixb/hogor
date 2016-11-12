package main

import (
	"log"
)

type Machine struct {
	state State
}

func NewMachine(initialState State) *Machine {
	m := Machine{initialState}
	return &m
}

func (m *Machine) Start() {
	log.Printf("Starting with '%s'", m.state.String())
	m.state.Enter(m)
}

// Deligate an event to current state
func (m *Machine) Event(pin uint, value uint) {
	log.Printf("Event: pin '%s' (%d), value %d", PinName(pin), pin, value)
	m.state.Event(m, pin, value)
}

// Transit to new state
func (m *Machine) Transit(newState State) {
	log.Printf("Transit from '%s' to '%s'", m.state.String(), newState.String())
	m.state.Leave(m)
	m.state = newState
	m.state.Enter(m)
}

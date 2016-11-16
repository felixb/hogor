package main

import (
	"testing"
)

// Enter does nothing
func TestOnState_Enter(t *testing.T) {
	m := MockMachine{}
	s := NewOnState()
	s.Enter(&m)

	m.AssertExpectations(t)
}

// Switch off transits to off state
func TestOnState_Event_switch_off(t *testing.T) {
	m := MockMachine{}
	s := NewOnState()
	m.On("Transit", STATE_OFF).Return(nil)

	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_OFF)

	m.AssertExpectations(t)
}

// Switch on transits nowhere
func TestOnState_Event_switch_on(t *testing.T) {
	m := MockMachine{}
	s := NewOnState()
	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_ON)

	m.AssertExpectations(t)
}

// Gate open transits to open state
func TestOnState_Event_gate_on(t *testing.T) {
	m := MockMachine{}
	s := NewOnState()
	m.On("Transit", STATE_OPEN).Return(nil)

	s.Event(&m, GPIO_GATE_PIN, GPIO_GATE_ON)

	m.AssertExpectations(t)
}

// Gate cloes transits nowhere
func TestOnState_Event_gate_off(t *testing.T) {
	m := MockMachine{}
	s := NewOnState()
	s.Event(&m, GPIO_GATE_PIN, GPIO_GATE_OFF)

	m.AssertExpectations(t)
}

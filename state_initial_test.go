package main

import (
	"testing"
)

// Transits to off state
func TestInitialState_Enter_off_off(t *testing.T) {
	m := MockMachine{}
	s := NewInititalState(GPIO_SWITCH_OFF, GPIO_GATE_OFF)

	m.On("Transit", STATE_OFF).Return(nil)
	s.Enter(&m)

	m.AssertExpectations(t)
}

// Transits to off state
func TestInitialState_Enter_off_on(t *testing.T) {
	m := MockMachine{}
	s := NewInititalState(GPIO_SWITCH_OFF, GPIO_GATE_ON)

	m.On("Transit", STATE_OFF).Return(nil)
	s.Enter(&m)

	m.AssertExpectations(t)
}

// Transits to waiting state
func TestInitialState_Enter_on_off(t *testing.T) {
	m := MockMachine{}
	s := NewInititalState(GPIO_SWITCH_ON, GPIO_GATE_OFF)

	m.On("Transit", STATE_WAITING).Return(nil)
	s.Enter(&m)

	m.AssertExpectations(t)
}

// Transits to waiting state
func TestInitialState_Enter_on_on(t *testing.T) {
	m := MockMachine{}
	s := NewInititalState(GPIO_SWITCH_ON, GPIO_GATE_ON)

	m.On("Transit", STATE_WAITING).Return(nil)
	s.Enter(&m)

	m.AssertExpectations(t)
}

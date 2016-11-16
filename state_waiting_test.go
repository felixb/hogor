package main

import (
	"testing"
)

// Enter does nothing
func TestWaitingState_Enter(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState()
	s.Enter(&m)

	m.AssertExpectations(t)
}

// Switch off transits to off state
func TestWaitingState_Event_switch_off(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState()
	m.On("Transit", STATE_OFF).Return(nil)

	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_OFF)

	m.AssertExpectations(t)
}

// Switch on transits nowhere
func TestWaitingState_Event_switch_On(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState()
	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_ON)

	m.AssertExpectations(t)
}
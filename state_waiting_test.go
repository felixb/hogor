package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Enter does nothing
func TestWaitingState_Enter(t *testing.T) {
	s0 := NewWaitingState()
	s1 := s0.Enter()

	// TODO mock timer

	assert.Nil(t, s1)
}

// Switch off transits to off state
func TestWaitingState_Event_switch_off(t *testing.T) {
	s0 := NewWaitingState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOff)

	assert.NotNil(t, s1)
	assert.Equal(t, "Off", s1.String())
}

// Switch on transits nowhere
func TestWaitingState_Event_switch_On(t *testing.T) {
	s0 := NewWaitingState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOn)

	assert.Nil(t, s1)
}

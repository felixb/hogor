package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Enter does nothing
func TestOpenState_Enter(t *testing.T) {
	s0 := NewOpenState()
	s1 := s0.Enter()

	assert.Nil(t, s1)
}

// Switch off transits to off state
func TestOpenState_Event_switch_off(t *testing.T) {
	s0 := NewOpenState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOff)

	assert.NotNil(t, s1)
	assert.Equal(t, "Off", s1.String())
}

// Switch on transits nowhere
func TestOpenState_Event_switch_On(t *testing.T) {
	s0 := NewOpenState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOn)

	assert.Nil(t, s1)
}

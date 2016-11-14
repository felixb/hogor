package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Enter does nothing
func TestOnState_Enter(t *testing.T) {
	s0 := NewOnState()
	s1 := s0.Enter()

	assert.Nil(t, s1)
}

// Switch off transits to off state
func TestOnState_Event_switch_off(t *testing.T) {
	s0 := NewOnState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOff)

	assert.NotNil(t, s1)
	assert.Equal(t, "Off", s1.String())
}

// Switch on transits nowhere
func TestOnState_Event_switch_on(t *testing.T) {
	s0 := NewOnState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOn)

	assert.Nil(t, s1)
}

// Gate open transits to open state
func TestOnState_Event_gate_on(t *testing.T) {
	s0 := NewOnState()
	s1 := s0.Event(gpioGate, gpioGateOn)

	assert.NotNil(t, s1)
	assert.Equal(t, "Open", s1.String())
}

// Gate cloes transits nowhere
func TestOnState_Event_gate_off(t *testing.T) {
	s0 := NewOnState()
	s1 := s0.Event(gpioGate, gpioGateOff)

	assert.Nil(t, s1)
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Enter does nothing
func TestOffState_Enter(t *testing.T) {
	s0 := NewOffState()
	s1 := s0.Enter()

	assert.Nil(t, s1)
}

// Switch on transits to waiting state
func TestOffState_Event_switch_on(t *testing.T) {
	s0 := NewOffState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOn)

	assert.NotNil(t, s1)
	assert.Equal(t, "Waiting", s1.String())
}

// Switch off transits nowhere
func TestOffState_Event_switch_off(t *testing.T) {
	s0 := NewOffState()
	s1 := s0.Event(gpioSwitch, gpioSwitchOff)

	assert.Nil(t, s1)
}

// Gate transits nowhere
func TestOffState_Event_gate(t *testing.T) {
	s0 := NewOffState()

	assert.Nil(t,  s0.Event(gpioGate, 0))
	assert.Nil(t,  s0.Event(gpioGate, 1))
}

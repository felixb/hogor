package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Transits to off state
func TestInitialState_Enter_off_off(t *testing.T) {
	s0 := NewInititalState(gpioSwitchOff, gpioGateOff)
	s1 := s0.Enter()

	assert.NotNil(t, s1)
	assert.Equal(t, "Off", s1.String())
}

// Transits to off state
func TestInitialState_Enter_off_on(t *testing.T) {
	s0 := NewInititalState(gpioSwitchOff, gpioGateOn)
	s1 := s0.Enter()

	assert.NotNil(t, s1)
	assert.Equal(t, "Off", s1.String())
}

// Transits to waiting state
func TestInitialState_Enter_on_off(t *testing.T) {
	s0 := NewInititalState(gpioSwitchOn, gpioGateOff)
	s1 := s0.Enter()

	assert.NotNil(t, s1)
	assert.Equal(t, "Waiting", s1.String())
}

// Transits to waiting state
func TestInitialState_Enter_on_on(t *testing.T) {
	s0 := NewInititalState(gpioSwitchOn, gpioGateOn)
	s1 := s0.Enter()

	assert.NotNil(t, s1)
	assert.Equal(t, "Waiting", s1.String())
}

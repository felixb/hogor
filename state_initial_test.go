package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Transits to off state
func TestInitialState_Enter_off_off(t *testing.T) {
	s := NewInititalState(gpioSwitchOff, gpioGateOff)
	m := NewMachine(s)
	m.Start()

	assert.Equal(t, "OffState", m.state.String())
}

// Transits to off state
func TestInitialState_Enter_off_on(t *testing.T) {
	s := NewInititalState(gpioSwitchOff, gpioGateOn)
	m := NewMachine(s)
	m.Start()

	assert.Equal(t, "OffState", m.state.String())
}

// Transits to waiting state
func TestInitialState_Enter_on_off(t *testing.T) {
	s := NewInititalState(gpioSwitchOn, gpioGateOff)
	m := NewMachine(s)
	m.Start()

	assert.Equal(t, "WaitingState", m.state.String())
}

// Transits to waiting state
func TestInitialState_Enter_on_on(t *testing.T) {
	s := NewInititalState(gpioSwitchOn, gpioGateOn)
	m := NewMachine(s)
	m.Start()

	assert.Equal(t, "WaitingState", m.state.String())
}

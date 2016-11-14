package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Enter does nothing
func TestWaitingState_Enter(t *testing.T) {
	s0 := NewWaitingState()
	s1 := s0.Enter()

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

func TestWaitingState_Tick(t *testing.T) {
	s := NewWaitingState()

	assert.Nil(t, s.Tick(time.Second))
	assert.Nil(t, s.Tick(2 * time.Second))
	assert.NotNil(t, s.Tick(65 * time.Second))
	assert.Equal(t, "On", s.Tick(65 * time.Second).String())
}
package main

import (
	"testing"
	"time"
)

// Enter does nothing
func TestWaitingState_Enter(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState(time.Minute)
	s.Enter(&m)

	m.AssertExpectations(t)
}

// Transits to on stage after 1 second
func TestWaitingState_Enter_timer(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState(time.Second)
	s.Enter(&m)
	// explicitly allowing call AFTER running Enter()
	m.On("Transit", STATE_ON).Return(nil)

	time.Sleep(time.Second * 2)

	m.AssertExpectations(t)
}

// Does not transit anywhere after leaving
func TestWaitingState_Leave(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState(time.Second)
	s.Enter(&m)
	s.Leave(&m)

	time.Sleep(time.Second * 2)

	m.AssertExpectations(t)
}

// Switch off transits to off state
func TestWaitingState_Event_switch_off(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState(time.Minute)
	m.On("Transit", STATE_OFF).Return(nil)

	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_OFF)

	m.AssertExpectations(t)
}

// Switch on transits nowhere
func TestWaitingState_Event_switch_On(t *testing.T) {
	m := MockMachine{}
	s := NewWaitingState(time.Minute)
	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_ON)

	m.AssertExpectations(t)
}

package main

import (
	"testing"
	"time"
)

// Enter does nothing
func TestWaitingState_Enter(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewWaitingState(&o, time.Minute, time.Minute)
	s.Enter(&m)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

// Transits to on stage after 2 * maxWait
func TestWaitingState_Enter_timer(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewWaitingState(&o, time.Minute, time.Millisecond * 100)
	s.Enter(&m)
	// explicitly allowing call AFTER running Enter()
	m.On("Transit", STATE_ON).Return(nil)

	time.Sleep(time.Millisecond * 200)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

// Does not transit anywhere after leaving
func TestWaitingState_Leave(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewWaitingState(&o, time.Minute, time.Millisecond * 100)
	s.Enter(&m)
	s.Leave(&m)

	time.Sleep(time.Millisecond * 200)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

// Blinks status LED and stops after leaving
func TestWaitingState_blink(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewWaitingState(&o, time.Millisecond * 10, time.Minute)
	o.On("High").Return(nil)
	o.On("Low").Return(nil)

	s.Enter(&m)
	time.Sleep(time.Millisecond * 100)
	s.Leave(&m)
	time.Sleep(time.Millisecond * 100)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
	o.AssertNumberOfCalls(t, "High", 5)
	o.AssertNumberOfCalls(t, "Low", 5)
}

// Switch off transits to off state
func TestWaitingState_Event_switch_off(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewWaitingState(&o, time.Minute, time.Minute)
	m.On("Transit", STATE_OFF).Return(nil)

	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_OFF)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

// Switch on transits nowhere
func TestWaitingState_Event_switch_On(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewWaitingState(&o, time.Minute, time.Minute)
	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_ON)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

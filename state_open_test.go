package main

import (
	"testing"
	"time"
)

// Enter does nothing
func TestOpenState_Enter(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOpenState(&os, &ob, time.Second)


	os.On("High").Return(nil)
	ob.On("High").Return(nil)

	s.Enter(&m)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)

	ob.On("Low").Return(nil)
	s.Leave(&m) // quit timers and stuff
}

// Transits to on stage after 1 second
func TestOpenState_Enter_timer(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOpenState(&os, &ob, time.Millisecond * 100)

	os.On("High").Return(nil)
	ob.On("High").Return(nil)

	s.Enter(&m)
	// explicitly allowing call AFTER running Enter()
	m.On("Transit", STATE_ON).Return(nil)

	time.Sleep(time.Millisecond * 200)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)

	ob.On("Low").Return(nil)
	s.Leave(&m) // quit timers and stuff
}

// Does not transit anywhere after leaving
func TestOpenState_Leave(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOpenState(&os, &ob, time.Millisecond * 100)

	os.On("High").Return(nil)
	ob.On("High").Return(nil)
	ob.On("Low").Return(nil)

	s.Enter(&m)
	s.Leave(&m)

	time.Sleep(time.Millisecond * 200)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}


// Switch off transits to off state
func TestOpenState_Event_switch_off(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOpenState(&os, &ob, time.Second)
	m.On("Transit", STATE_OFF).Return(nil)

	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_OFF)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

// Switch on transits nowhere
func TestOpenState_Event_switch_on(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOpenState(&os, &ob, time.Second)
	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_ON)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

// Switch on transits to on state
func TestOpenState_Event_gate_off(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOpenState(&os, &ob, time.Second)
	m.On("Transit", STATE_ON).Return(nil)

	s.Event(&m, GPIO_GATE_PIN, GPIO_GATE_OFF)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

// Switch on transits nowhere
func TestOpenState_Event_gate_on(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOpenState(&os, &ob, time.Second)
	s.Event(&m, GPIO_GATE_PIN, GPIO_GATE_ON)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

package main

import (
	"testing"
	"time"
)

// Enter does nothing
func TestOnState_Enter(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOnState(&os, &ob, time.Second)
	os.On("High").Return(nil)

	s.Enter(&m)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

// Switch off transits to off state
func TestOnState_Event_switch_off(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOnState(&os, &ob, time.Second)
	m.On("Transit", STATE_OFF).Return(nil)

	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_OFF)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

// Switch on transits nowhere
func TestOnState_Event_switch_on(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOnState(&os, &ob, time.Second)
	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_ON)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

// Gate closes transits nowhere
func TestOnState_Event_gate_off(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOnState(&os, &ob, time.Second)
	s.Event(&m, GPIO_GATE_PIN, GPIO_GATE_OFF)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}


// Gate open pulls down and up the bell
func TestOnState_Event_gate_on(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOnState(&os, &ob, time.Millisecond * 100)

	ob.On("Low").Return(nil)
	ob.On("High").Return(nil)

	s.Event(&m, GPIO_GATE_PIN, GPIO_GATE_ON)
	time.Sleep(time.Millisecond * 200)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}

// Pull up bell
func TestOnState_Leave(t *testing.T) {
	m := MockMachine{}
	os := MockOutput{}
	ob := MockOutput{}
	s := NewOnState(&os, &ob, time.Second)

	ob.On("High").Return(nil)

	s.Leave(&m)

	m.AssertExpectations(t)
	os.AssertExpectations(t)
	ob.AssertExpectations(t)
}
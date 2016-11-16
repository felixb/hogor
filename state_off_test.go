package main

import (
	"testing"
)

// Enter does nothing but setting status led to low
func TestOffState_Enter(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewOffState(&o)
	o.On("Low").Return(nil)

	s.Enter(&m)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

// Switch on transits to waiting state
func TestOffState_Event_switch_on(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewOffState(&o)
	m.On("Transit", STATE_WAITING).Return(nil)

	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_ON)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

// Switch off transits nowhere
func TestOffState_Event_switch_off(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewOffState(&o)
	s.Event(&m, GPIO_SWITCH_PIN, GPIO_SWITCH_OFF)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

// Gate transits nowhere
func TestOffState_Event_gate(t *testing.T) {
	m := MockMachine{}
	o := MockOutput{}
	s := NewOffState(&o)
	s.Event(&m, GPIO_GATE_PIN, 0)
	s.Event(&m, GPIO_GATE_PIN, 1)

	m.AssertExpectations(t)
	o.AssertExpectations(t)
}

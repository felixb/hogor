package main

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockState struct {
	mock.Mock
}

func (s *MockState) Enter(m *Machine)                       { s.Called(m) }
func (s *MockState) Leave(m *Machine)                       { s.Called(m) }
func (s *MockState) Event(m *Machine, pin uint, value uint) { s.Called(m, pin, value) }
func (s *MockState) String() string                         { return "MockState" }

// Machine.Start calls Enter on initial state
func TestMachine_Start(t *testing.T) {
	s := new(MockState)
	m := NewMachine(s)

	s.On("Enter", m).Return()

	m.Start()

	s.AssertExpectations(t)
}

// Machine.Event delegates to Machine.state.Event
func TestMachine_Event(t *testing.T) {
	var pin uint
	var value uint
	pin = 1
	value = 2

	s := new(MockState)
	m := NewMachine(s)

	s.On("Event", m, pin, value).Return()

	m.Event(pin, value)

	s.AssertExpectations(t)
}

// Machine.Transit calls Leave on old state, Enter on new state
func TestMachine_Transit(t *testing.T) {
	s0 := new(MockState)
	m := NewMachine(s0)

	s0.On("Leave", m).Return()

	s1 := new(MockState)
	s1.On("Enter", m).Return()

	m.Transit(s1)

	s0.AssertExpectations(t)
}

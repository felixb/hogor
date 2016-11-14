package main

import (
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"github.com/stretchr/testify/assert"
)

type MockState struct {
	mock.Mock
}

func (s *MockState) Enter() State {
	args := s.Called()

	if r := args.Get(0); r != nil {
		return args.Get(0).(State)
	} else {
		return nil
	}
}
func (s *MockState) Leave() {
	s.Called()
}
func (s *MockState) Event(pin uint, value uint) State {
	args := s.Called(pin, value)

	if r := args.Get(0); r != nil {
		return args.Get(0).(State)
	} else {
		return nil
	}
}
func (s *MockState) String() string {
	return "MockState"
}

// Machine.Start calls Enter on initial state
func TestMachine_Start_nil(t *testing.T) {
	s := new(MockState)
	m := NewMachine(s)

	s.On("Enter").Return(nil)

	m.Start()

	s.AssertExpectations(t)
}

// Machine.Start calls Enter and moves to that state
func TestMachine_Start_state(t *testing.T) {
	s0 := new(MockState)
	s1 := new(MockState)

	m := NewMachine(s0)

	s0.On("Enter").Return(s1)
	s1.On("Enter").Return(nil)

	m.Start()

	s0.AssertExpectations(t)
	s1.AssertExpectations(t)
	assert.Equal(t, s1, m.state)
}

// Machine.Event delegates to Machine.state.Event
func TestMachine_Event_nil(t *testing.T) {
	var pin uint
	var value uint
	pin = 1
	value = 2

	s := new(MockState)
	m := NewMachine(s)

	s.On("Event", pin, value).Return(nil)

	m.Event(pin, value)

	s.AssertExpectations(t)
}

// Machine.Event delegates to Machine.state.Event
func TestMachine_Event_state(t *testing.T) {
	var pin uint
	var value uint
	pin = 1
	value = 2

	s0 := new(MockState)
	s1 := new(MockState)

	m := NewMachine(s0)

	s0.On("Event", pin, value).Return(s1)
	s1.On("Enter").Return(nil)

	m.Event(pin, value)

	s0.AssertExpectations(t)
	s1.AssertExpectations(t)
	assert.Equal(t, s1, m.state)
}

// Machine.Transit calls Enter on new state
func TestMachine_Transit(t *testing.T) {
	s0 := new(MockState)
	m := NewMachine(s0)

	s1 := new(MockState)
	s1.On("Enter").Return(nil)

	m.Transit(s1)

	s0.AssertExpectations(t)
	s1.AssertExpectations(t)
	assert.Equal(t, s1, m.state)
}

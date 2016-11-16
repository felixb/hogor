package main

import (
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"github.com/stretchr/testify/assert"
)

type MockMachine struct {
	mock.Mock
}

func (m *MockMachine) AddState(id StateId, state State) error {
	args := m.Called(id, state)
	return args.Error(0)
}

func (m *MockMachine) CurrentState() State {
	args := m.Called()
	return args[0].(State)
}

func (m *MockMachine) State(id StateId) State {
	args := m.Called(id)
	return args[0].(State)
}
func (m *MockMachine) Start(id StateId) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMachine) Transit(id StateId) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMachine) Event(pin uint, value uint) {
	m.Called(pin, value)
}

type MockState struct {
	mock.Mock
}

func (s *MockState) Enter(m StateMachine) {
	s.Called(m)
}
func (s *MockState) Leave(m StateMachine) {
	s.Called(m)
}
func (s *MockState) Event(m StateMachine, pin uint, value uint) {
	s.Called(m, pin, value)
}

func (s *MockState) String() string {
	return "MockState"
}

func TestMachine_AddState(t *testing.T) {
	s := new(MockState)
	m := NewMachine()

	assert.Nil(t, m.State(1))
	assert.Nil(t, m.AddState(1, s))
	assert.Equal(t, s, m.State(1))
	assert.Error(t, m.AddState(1, s))
}

// Machine.Start calls Enter on initial state
func TestMachine_Start(t *testing.T) {
	s := new(MockState)
	m := NewMachine()
	m.AddState(0, s)

	s.On("Enter", m).Return()

	m.Start(0)

	s.AssertExpectations(t)
}

func TestMachine_Start_missing_state(t *testing.T) {
	m := NewMachine()
	assert.Error(t, m.Start(0))
}

// Machine.Event delegates to Machine.state.Event
func TestMachine_Event(t *testing.T) {
	var pin uint
	var value uint
	pin = 1
	value = 2

	s := new(MockState)
	m := NewMachine()
	m.AddState(0, s)

	s.On("Enter", m).Return()
	s.On("Event", m, pin, value).Return()

	m.Start(0)
	m.Event(pin, value)

	s.AssertExpectations(t)
}

// Machine.Transit calls Enter on new state
func TestMachine_Transit(t *testing.T) {
	var id0, id1 StateId
	id0 = 0
	id1 = 1

	s0 := new(MockState)
	s1 := new(MockState)
	m := NewMachine()
	m.AddState(id0, s0)
	m.AddState(id1, s1)

	s0.On("Enter", m).Return()
	s0.On("Leave", m).Return()
	s1.On("Enter", m).Return()

	m.Start(0)
	m.Transit(1)

	s0.AssertExpectations(t)
	s1.AssertExpectations(t)
	assert.Equal(t, id1, m.currentState)
	assert.Equal(t, s1, m.CurrentState())
}

func TestMachine_Transit_missing_state(t *testing.T) {
	s0 := new(MockState)
	m := NewMachine()
	m.AddState(0, s0)

	assert.Error(t, m.Transit(1))
}

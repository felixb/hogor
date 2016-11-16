package main

import (
	"github.com/stretchr/testify/mock"
)

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

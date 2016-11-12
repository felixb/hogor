package main

type WaitingState struct{}

func NewWaitingState() State {
	s := WaitingState{}
	return &s
}

func (s *WaitingState) Enter(m *Machine) {}

func (s *WaitingState) Leave(m *Machine) {}

func (s *WaitingState) Event(m *Machine, pin uint, value uint) {}

func (s *WaitingState) String() string {
	return "WaitingState"
}

package main

type OffState struct{}

func NewOffState() State {
	s := OffState{}
	return &s
}

func (s *OffState) Enter(m *Machine) {}

func (s *OffState) Leave(m *Machine) {}

func (s *OffState) Event(m *Machine, pin uint, value uint) {}

func (s *OffState) String() string {
	return "OffState"
}

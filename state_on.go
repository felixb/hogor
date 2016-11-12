package main

type OnState struct{}

func NewOnState() State {
	s := OnState{}
	return &s
}

func (s *OnState) Enter(m *Machine) {}

func (s *OnState) Leave(m *Machine) {}

func (s *OnState) Event(m *Machine, pin uint, value uint) {}

func (s *OnState) String() string {
	return "OnState"
}

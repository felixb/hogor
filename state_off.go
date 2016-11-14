package main

type OffState struct{}

func NewOffState() State {
	s := OffState{}
	return &s
}

func (s *OffState) Enter() State {
	return nil
}

func (s *OffState) Leave() {}

func (s *OffState) Event(pin uint, value uint) State {
	return nil
}

func (s *OffState) String() string {
	return "Off"
}

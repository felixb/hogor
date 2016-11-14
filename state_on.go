package main

type OnState struct{}

func NewOnState() State {
	s := OnState{}
	return &s
}
func (s *OnState) Enter() State {
	return nil
}

func (s *OnState) Leave() {}

func (s *OnState) Event(pin uint, value uint) State {
	return nil
}

func (s *OnState) String() string {
	return "On"
}

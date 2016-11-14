package main

type WaitingState struct{}

func NewWaitingState() State {
	s := WaitingState{}
	return &s
}

func (s *WaitingState) Enter() State {
	return nil
}

func (s *WaitingState) Leave() {}

func (s *WaitingState) Event(pin uint, value uint) State {
	return nil
}

func (s *WaitingState) String() string {
	return "Waiting"
}

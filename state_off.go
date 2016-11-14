package main

type OffState struct{}

func NewOffState() State {
	s := OffState{}
	return &s
}

func (s *OffState) Enter() State {
	// TODO status led -> off
	return nil
}

func (s *OffState) Leave() {}

func (s *OffState) Event(pin uint, value uint) State {
	if pin == gpioSwitch && value == gpioSwitchOn {
		return NewWaitingState()
	} else {
		return nil
	}
}

func (s *OffState) String() string {
	return "Off"
}

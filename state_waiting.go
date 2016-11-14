package main

type WaitingState struct{}

func NewWaitingState() State {
	s := WaitingState{}
	return &s
}

func (s *WaitingState) Enter() State {
	// TODO schedule timer to transit to on state
	// TODO blink status led
	return nil
}

func (s *WaitingState) Event(pin uint, value uint) State {
	if pin == gpioSwitch && value == gpioSwitchOff {
		return NewOffState()
	} else {
		return nil
	}
}

func (s *WaitingState) String() string {
	return "Waiting"
}

package main

type OnState struct{}

func NewOnState() State {
	s := OnState{}
	return &s
}
func (s *OnState) Enter() State {
	// TODO status led -> on
	return nil
}

func (s *OnState) Leave() {}

func (s *OnState) Event(pin uint, value uint) State {
	if pin == gpioSwitch && value == gpioSwitchOff {
		return NewOffState()
	}

	if pin == gpioGate && value == gpioGateOn {
		return NewOpenState()
	}

	return nil
}

func (s *OnState) String() string {
	return "On"
}

package main

type OpenState struct{}

func NewOpenState() State {
	s := OpenState{}
	return &s
}
func (s *OpenState) Enter() State {
	// TODO schedule transit to on state
	// TODO ring the bell
	return nil
}

func (s *OpenState) Event(pin uint, value uint) State {
	if pin == gpioSwitch && value == gpioSwitchOff {
		return NewOffState()
	} else {
		return nil
	}
}

func (s *OpenState) String() string {
	return "Open"
}

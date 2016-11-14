package main

import (
	"time"
)

type OpenState struct{}

func NewOpenState() *OpenState {
	s := OpenState{}
	return &s
}
func (s *OpenState) Enter() State {
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

func (s *OpenState) Tick(d time.Duration) State {
	if d < max_wait {
		return nil
	} else {
		return NewOnState()
	}
}

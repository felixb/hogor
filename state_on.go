package main

import (
	"time"
)

type OnState struct {
	statusLed    Output
	bell         Output
	bellDuration time.Duration
	timer        *time.Timer
}

func NewOnState(statusLed Output, bell Output, bellDuration time.Duration) State {
	s := OnState{
		statusLed: statusLed,
		bell: bell,
		bellDuration: bellDuration,
	}
	return &s
}
func (s *OnState) Enter(m StateMachine) {
	s.statusLed.High()
}

func (s *OnState) Event(m StateMachine, pin uint, value uint) {
	if pin == GPIO_SWITCH_PIN && value == GPIO_SWITCH_OFF {
		m.Transit(STATE_OFF)
	} else if pin == GPIO_GATE_PIN && value == GPIO_GATE_ON {
        s.bell.Low()
		s.timer = time.AfterFunc(s.bellDuration, func() {
		    s.bell.High()
			s.timer = nil
		})
	}
}

func (s *OnState) Leave(m StateMachine) {
	if s.timer != nil {
		s.timer.Stop()
		s.timer = nil
	}
	s.bell.High()
}

func (s *OnState) String() string {
	return "On"
}

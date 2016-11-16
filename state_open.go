package main

import (
	"time"
)

type OpenState struct {
	statusLed Output
	bell      Output
	maxWait   time.Duration
	timer     *time.Timer
	inStep    bool
}

func NewOpenState(statusLed Output, bell Output, maxWait time.Duration) *OpenState {
	s := OpenState{
		statusLed: statusLed,
		bell: bell,
		maxWait: maxWait,
	}
	return &s
}
func (s *OpenState) Enter(m StateMachine) {
	s.inStep = true
	s.statusLed.High()
	s.bell.High()
	s.startTimer(m)
}

func (s *OpenState) Event(m StateMachine, pin uint, value uint) {
	if pin == GPIO_SWITCH_PIN && value == GPIO_SWITCH_OFF {
		m.Transit(STATE_OFF)
	} else if pin == GPIO_GATE_PIN && value == GPIO_GATE_OFF {
		m.Transit(STATE_ON)
	}
}

func (s *OpenState) Leave(m StateMachine) {
	s.inStep = false
	s.bell.Low()
	s.stopTimer()
}

func (s *OpenState) String() string {
	return "Open"
}

func (s *OpenState) startTimer(m StateMachine) {
	s.timer = time.AfterFunc(s.maxWait, func() {
		if s.inStep {
			m.Transit(STATE_ON)
		}
	})
}

func (s *OpenState) stopTimer() {
	s.timer.Stop()
}

package main

import (
	"time"
)

type WaitingState struct {
	statusLed      Output
	statusValue    bool
	statusInterval time.Duration
	maxWait        time.Duration
	timer          *time.Timer
	ticker         *time.Ticker
	inStep         bool
}

func NewWaitingState(statusLed Output, statusInterval time.Duration, maxWait time.Duration) *WaitingState {
	s := WaitingState{
		statusLed: statusLed,
		statusValue: false,
		statusInterval: statusInterval,
		maxWait: maxWait,
	}
	return &s
}

func (s *WaitingState) Enter(m StateMachine) {
	s.inStep = true
	s.startTimer(m)
	s.startTicker(m)
}

func (s *WaitingState) Leave(m StateMachine) {
	s.inStep = false
	s.stopTimer()
	s.stopTicker()
}

func (s *WaitingState) Event(m StateMachine, pin uint, value uint) {
	if pin == GPIO_SWITCH_PIN && value == GPIO_SWITCH_OFF {
		m.Transit(STATE_OFF)
	}
}

func (s *WaitingState) String() string {
	return "Waiting"
}

func (s *WaitingState) startTimer(m StateMachine) {
	s.timer = time.AfterFunc(s.maxWait, func() {
		if s.inStep {
			m.Transit(STATE_ON)
		}
	})
}

func (s *WaitingState) stopTimer() {
	s.timer.Stop()
}

func (s *WaitingState) startTicker(m StateMachine) {
	s.ticker = time.NewTicker(s.statusInterval)
	go func() {
		for range s.ticker.C {
			if s.statusValue {
				s.statusLed.Low()
			} else {
				s.statusLed.High()
			}
			s.statusValue = !s.statusValue
		}
	}()
}

func (s *WaitingState) stopTicker() {
	s.ticker.Stop()
}
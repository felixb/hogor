package main

import (
	"time"
)

type WaitingState struct {
	max_wait      time.Duration
	timer         *time.Timer
	stopTimerChan chan bool
}

func NewWaitingState(max_wait time.Duration) *WaitingState {
	s := WaitingState{max_wait: max_wait}
	return &s
}

func (s *WaitingState) Enter(m StateMachine) {
	// TODO blink status led
	s.startTimer(m)
}

func (s *WaitingState) Leave(m StateMachine) {
	s.stopTimer()
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
	s.stopTimerChan = make(chan bool)
	s.timer = time.NewTimer(s.max_wait)
	go func() {
		select {
		case <-s.timer.C:
			m.Transit(STATE_ON)
		case <-s.stopTimerChan:
			return
		}
	}()
}

func (s *WaitingState) stopTimer() {
	s.stopTimerChan <- true
	if !s.timer.Stop() {
		<-s.timer.C
	}
}
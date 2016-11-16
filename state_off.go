package main

type OffState struct {
	statusLed Output
}

func NewOffState(statusLed Output) State {
	s := OffState{statusLed: statusLed}
	return &s
}

func (s *OffState) Enter(m StateMachine) {
	s.statusLed.Low()
}

func (s *OffState) Event(m StateMachine, pin uint, value uint) {
	if pin == GPIO_SWITCH_PIN && value == GPIO_SWITCH_ON {
		m.Transit(STATE_WAITING)
	}
}

func (s *OffState) Leave(m StateMachine) {}

func (s *OffState) String() string {
	return "Off"
}

package main

type OnState struct {
	statusLed Output
}

func NewOnState(statusLed Output) State {
	s := OnState{statusLed: statusLed}
	return &s
}
func (s *OnState) Enter(m StateMachine) {
	s.statusLed.High()
}

func (s *OnState) Event(m StateMachine, pin uint, value uint) {
	if pin == GPIO_SWITCH_PIN && value == GPIO_SWITCH_OFF {
		m.Transit(STATE_OFF)
	} else if pin == GPIO_GATE_PIN && value == GPIO_GATE_ON {
		m.Transit(STATE_OPEN)
	}
}

func (s *OnState) Leave(m StateMachine) {}

func (s *OnState) String() string {
	return "On"
}

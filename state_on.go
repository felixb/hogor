package main

type OnState struct{}

func NewOnState() State {
	s := OnState{}
	return &s
}
func (s *OnState) Enter(m StateMachine) {
	// TODO status led -> on
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

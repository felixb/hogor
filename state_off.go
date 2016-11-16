package main

type OffState struct{}

func NewOffState() State {
	s := OffState{}
	return &s
}

func (s *OffState) Enter(m StateMachine) {
	// TODO status led -> off
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

package main

type State interface {
	Enter(m StateMachine)
	Leave(m StateMachine)
	Event(m StateMachine, pin uint, value uint)
	String() string
}

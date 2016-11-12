package main

type State interface {
	Enter(m *Machine)
	Leave(m *Machine)
	Event(m *Machine, pin uint, value uint)
	String() string
}

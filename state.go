package main

type State interface {
	Enter() State
	Leave()
	Event(pin uint, value uint) State
	String() string
}

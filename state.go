package main

import (
	"time"
)

type State interface {
	Enter() State
	Event(pin uint, value uint) State
	String() string
}

type Ticker interface {
	Tick(time.Duration) State
}

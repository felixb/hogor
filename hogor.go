package main

import (
	"github.com/brian-armstrong/gpio"
	"log"
)

const (
	GPIO_SWITCH_PIN = 2
	GPIO_SWITCH_OFF = 1
	GPIO_SWITCH_ON = 0
	GPIO_GATE_PIN = 3
	GPIO_GATE_OFF = 1
	GPIO_GATE_ON = 0
	// GPIO_LIGHT_PIN  = 4
	// GPIO_LIGHT_BELL = 5

	STATE_INITIAL StateId = 0
	STATE_OFF StateId = 1
	STATE_ON StateId = 2
	STATE_OPEN StateId = 3
	STATE_WAITING StateId = 4
)

var (
	pinNames = map[uint]string{
		2: "switch",
		3: "gate",
		4: "light",
		5: "bell",
	}
)

func PinName(pin uint) string {
	n := pinNames[pin]
	if n != "" {
		return n
	} else {
		return "unknown pin"
	}
}

func readIgnoreErrors(p gpio.Pin) uint {
	if v, err := p.Read(); err != nil {
		log.Printf("Error reading pin '%s' (%d)", PinName(p.Number), p.Number)
		return v
	} else {
		return v
	}
}

func check(err error) {
	if err != nil {
		log.Panicf("Unexpected error: %s", err)
	}
}

func main() {
	pinSwitch := gpio.NewInput(GPIO_SWITCH_PIN)
	pinGate := gpio.NewInput(GPIO_GATE_PIN)

	watcher := gpio.NewWatcher()
	watcher.AddPin(GPIO_SWITCH_PIN)
	watcher.AddPin(GPIO_GATE_PIN)
	defer watcher.Close()
	defer pinSwitch.Close()

	m := NewMachine()
	check(m.AddState(STATE_INITIAL, NewInititalState(readIgnoreErrors(pinSwitch), readIgnoreErrors(pinGate))))
	check(m.AddState(STATE_OFF, NewOffState()))
	check(m.AddState(STATE_ON, NewOnState()))
	check(m.AddState(STATE_OPEN, NewOpenState()))
	check(m.AddState(STATE_WAITING, NewWaitingState()))

	check(m.Start(STATE_INITIAL))

	for {
		m.Event(watcher.Watch())
	}
}

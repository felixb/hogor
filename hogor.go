package main

import (
	"github.com/brian-armstrong/gpio"
	"log"
)

const (
	gpioSwitch    = 2
	gpioSwitchOff = 1
	gpioSwitchOn  = 0
	gpioGate      = 3
	gpioGateOff   = 1
	gpioGateOn    = 0
	// gpioLight  = 4
	// gpioBell   = 5
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

func main() {
	pinSwitch := gpio.NewInput(gpioSwitch)
	pinGate := gpio.NewInput(gpioGate)

	watcher := gpio.NewWatcher()
	watcher.AddPin(gpioSwitch)
	watcher.AddPin(gpioGate)
	defer watcher.Close()
	defer pinSwitch.Close()

	s := NewInititalState(readIgnoreErrors(pinSwitch), readIgnoreErrors(pinGate))
	m := NewMachine(s)
	m.Start()

	for {
		m.Event(watcher.Watch())
	}
}

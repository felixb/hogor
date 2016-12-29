package main

import (
	"github.com/brian-armstrong/gpio"
	"log"
	"time"
)

const (
	GPIO_SWITCH_PIN = 2
	GPIO_SWITCH_OFF = 1
	GPIO_SWITCH_ON = 0
	GPIO_GATE_PIN = 3
	GPIO_GATE_OFF = 1
	GPIO_GATE_ON = 0
	GPIO_LIGHT_PIN = 4
	GPIO_BELL_PIN = 5

	STATE_OFF StateId = 1
	STATE_ON StateId = 2
	STATE_OPEN StateId = 3
	STATE_WAITING StateId = 4

	WAITING_DURATION = time.Minute
	STATUS_BLINK_INTERVAL = time.Second
	OPEN_DURATION = time.Second * 5

	CHECK_INTERVAL = time.Second * 2
)

func check(err error) {
	if err != nil {
		log.Panicf("Unexpected error: %s", err)
	}
}

func main() {
	pinSwitch := gpio.NewInput(GPIO_SWITCH_PIN)
	defer pinSwitch.Close()
	pinGate := gpio.NewInput(GPIO_GATE_PIN)
	defer pinGate.Close()
	pinLight := gpio.NewOutput(GPIO_LIGHT_PIN, false)
	defer pinLight.Close()
	pinBell := gpio.NewOutput(GPIO_BELL_PIN, false)
	defer pinBell.Close()

	m := NewMachine()
	check(m.AddState(STATE_OFF, NewOffState(&pinLight)))
	check(m.AddState(STATE_ON, NewOnState(&pinLight)))
	check(m.AddState(STATE_OPEN, NewOpenState(&pinLight, &pinBell, OPEN_DURATION)))
	check(m.AddState(STATE_WAITING, NewWaitingState(&pinLight, STATUS_BLINK_INTERVAL, WAITING_DURATION)))

	check(m.Start(STATE_OFF))

	d := NewDriver(CHECK_INTERVAL, m.Event)
	defer d.Close()

	d.AddInput(GPIO_SWITCH_PIN, "switch", pinSwitch, GPIO_SWITCH_OFF)
	d.AddInput(GPIO_GATE_PIN, "gate", pinGate, GPIO_GATE_OFF)
	d.Start()
	d.Loop()
}

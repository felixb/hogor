package main

import (
	"time"
	"log"
)

type EventFunc func(pin uint, value uint)

type Driver struct {
	interval  time.Duration
	eventFunc EventFunc
	inputs    map[uint]Input
	names     map[uint]string
	state     map[uint]uint
	ticker    *time.Ticker
}

func NewDriver(interval time.Duration, eventFunc EventFunc) *Driver {
	d := Driver{
		interval: interval,
		eventFunc: eventFunc,
		inputs: make(map[uint]Input),
		names: make(map[uint]string),
		state: make(map[uint]uint),
	}
	return &d
}

func (d *Driver) Start() {
	d.ticker = time.NewTicker(d.interval)
}

func (d *Driver) Close() {
	if d.ticker != nil {
		d.ticker.Stop()
	}
}

func (d *Driver) AddInput(pin uint, name string, i Input, value uint) {
	d.names[pin] = name
	d.state[pin] = value
	d.inputs[pin] = i
}

func (d *Driver) readInputs() {
	for pin, i := range d.inputs {
		if v, err := i.Read(); err != nil {
			log.Printf("Error reading pin %s (%d): %s", d.names[pin], pin, err)
		} else {
			if v != d.state[pin] {
				log.Printf("Event: %s (%d) %d -> %d", d.names[pin], pin, d.state[pin], v)
				d.state[pin] = v
				d.eventFunc(pin, v)
			}
		}

	}
}

func (d *Driver) Loop() {
	for range d.ticker.C {
		d.readInputs()
	}
}

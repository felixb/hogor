package main

import (
	"log"
	"time"
)

const (
	tick = time.Second
)

type Machine struct {
	state  State
	ticker *time.Ticker
}

func NewMachine(initialState State) *Machine {
	m := Machine{
		state: initialState,
		ticker: nil,
	}
	return &m
}

func (m *Machine) Start() {
	log.Printf("Starting with '%s'", m.state.String())
	if s := m.state.Enter(); s != nil {
		m.Transit(s)
	}
}

// Delegate an event to current state
func (m *Machine) Event(pin uint, value uint) {
	log.Printf("Event: pin '%s' (%d), value %d", PinName(pin), pin, value)
	if s := m.state.Event(pin, value); s != nil {
		m.Transit(s)
	}
}

// Transit to new state
func (m *Machine) Transit(newState State) {
	log.Printf("Transit from '%s' to '%s'", m.state.String(), newState.String())
	m.state = newState
	m.state.Enter()

	if m.ticker != nil {
		m.stopTicker()
	}
	if t, ok := m.state.(Ticker); ok {
		m.startTicker(t)
	}
}

func (m *Machine) startTicker(t Ticker) {
	log.Printf("Starting ticker")
	m.ticker = time.NewTicker(tick)
	start := time.Now()
	go func() {
		for range m.ticker.C {
			if s := t.Tick(time.Since(start)); s != nil {
				m.Transit(s)
			}
		}
	}()
}

func (m *Machine) stopTicker() {
	log.Printf("Stopping ticker")
	m.ticker.Stop()
	m.ticker = nil
}

package main

import (
	"log"
	"time"
	"fmt"
)

const (
	tick = time.Second
)

type StateId uint

type StateMachine interface {
	AddState(id StateId, state State) error
	CurrentState() State
	State(id StateId) State
	Start(id StateId) error
	Transit(id StateId) error
	Event(pin uint, value uint)
}

type Machine struct {
	currentState StateId
	states       map[StateId]State
	ticker       *time.Ticker
}

func NewMachine() *Machine {
	m := Machine{
		states: make(map[StateId]State),
		ticker: nil,
	}
	return &m
}

func (m *Machine) AddState(id StateId, state State) error {
	if s, ok := m.states[id]; ok {
		return fmt.Errorf("State with id %d already exists: %s", id, s.String())
	}

	m.states[id] = state
	return nil
}

func (m *Machine) CurrentState() State {
	return m.State(m.currentState)
}

func (m *Machine) State(id StateId) State {
	return m.states[id]
}

func (m *Machine) Start(id StateId) error {
	if s := m.State(id); s == nil {
		return fmt.Errorf("State with id %d does not exist", id)
	} else {
		log.Printf("Starting with state id %d: %s", id, s.String())
		m.currentState = id
		s.Enter(m)
		return nil
	}
}

// Transit to new state
func (m *Machine) Transit(id StateId) error {
	if s := m.State(id); s == nil {
		return fmt.Errorf("State with id %d does not exist", id)
	} else {
		oldState := m.CurrentState()
		log.Printf("Transit from %d '%s' to %d '%s'", m.currentState, oldState.String(), id, s.String())
		oldState.Leave(m)
		m.currentState = id
		s.Enter(m)

		return nil
	}
}

// Delegate an event to current state
func (m *Machine) Event(pin uint, value uint) {
	log.Printf("Event: pin '%s' (%d), value %d", PinName(pin), pin, value)
	m.CurrentState().Event(m, pin, value)
}

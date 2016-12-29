package main

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestDriver_AddInput(t *testing.T) {
	i := MockInput{}
	d := NewDriver(time.Millisecond, nil)

	d.AddInput(0, "some-input", &i, 9)

	assert.Equal(t, "some-input", d.names[0])
	assert.Equal(t, uint(9), d.state[0])

	i.AssertExpectations(t)
}

func TestDriver_ReadInputs_wo_change(t *testing.T) {
	i := MockInput{}
	d := NewDriver(time.Millisecond, func(pin uint, value uint) {
		assert.Fail(t, "unexpected event")
	})

	i.On("Read").Return(uint(9), nil)

	d.AddInput(0, "some-input", &i, 9)
	d.readInputs()

	i.AssertExpectations(t)
}

func TestDriver_ReadInputs_w_change(t *testing.T) {
	i := MockInput{}
	called := false
	d := NewDriver(time.Millisecond, func(pin uint, value uint) {
		assert.False(t, called, "unexpected second call")
		called = true
		assert.Equal(t, uint(0), pin)
		assert.Equal(t, uint(10), value)
	})

	i.On("Read").Return(uint(10), nil)

	d.AddInput(0, "some-input", &i, 9)
	d.readInputs()
	d.readInputs()

	assert.True(t, called)
	i.AssertExpectations(t)
}

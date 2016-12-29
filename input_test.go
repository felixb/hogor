package main

import (
	"github.com/stretchr/testify/mock"
)

type MockInput struct {
	mock.Mock
}

func (i *MockInput) Read() (value uint, err error) {
	args := i.Called()
	return args.Get(0).(uint), args.Error(1)
}

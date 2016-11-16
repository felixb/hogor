package main

import (
	"github.com/stretchr/testify/mock"
)

type MockOutput struct {
	mock.Mock
}

func (o *MockOutput) High() error {
	args := o.Called()
	return args.Error(0)
}

func (o *MockOutput) Low() error {
	args := o.Called()
	return args.Error(0)
}
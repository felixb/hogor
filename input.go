package main

type Input interface {
	Read() (value uint, err error)
}

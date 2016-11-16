package main

type Output interface {
	High() error
	Low() error
}

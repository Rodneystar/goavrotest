package main

import (
	"github.com/linkedin/goavro"
)

type Nose struct {
	Size string
	Shape string
}

type Person struct {
	Name string
	Noses []Nose
}

func athing() {
	goavro.NewCodec()
}


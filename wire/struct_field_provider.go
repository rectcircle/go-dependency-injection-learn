package main

import "github.com/google/wire"

type Foo4 struct {
	S string
	N int
	F float64
}

func NewFoo4() Foo4 {
	return Foo4{S: "Hello, World!", N: 1, F: 3.14}
}

var StructFieldProviderSet = wire.NewSet(
	NewFoo4,
	wire.FieldsOf(new(Foo4), "S"))

package main

import (
	"io"
	"os"

	"github.com/google/wire"
)

type Foo3 struct {
	X int
}

var BindValueSet1 = wire.NewSet(wire.Value(Foo3{X: 42}))
var BindValueSet2 = wire.NewSet(wire.InterfaceValue(new(io.Reader), os.Stdin))

package main

import "github.com/google/wire"

type D1 struct {
	D2 *D2 `autowire:""`
}

type D2 struct {
	D1 *D1 `autowire:""`
}

var CircularDependencySet = wire.NewSet(
	wire.Struct(new(D1), "*"),
	wire.Struct(new(D2), "*"),
)

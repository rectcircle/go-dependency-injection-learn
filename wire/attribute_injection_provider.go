package main

import "github.com/google/wire"

type Foo2 int
type Bar2 int

func newFoo2() Foo2 { return 1 }

func newBar2() Bar2 { return 2 }

type FooBar2 struct {
	MyFoo2   Foo2
	MyBar2   Bar2
	MyBar2_2 Bar2 `wire:"-"` // 忽略该字段
}

var StructProviderSet = wire.NewSet(
	newFoo2,
	newBar2,
	wire.Struct(new(FooBar2), "MyFoo2")) // 只绑定一个参数

var StructProviderSet2 = wire.NewSet(
	newFoo2,
	newBar2,
	wire.Struct(new(FooBar2), "*")) // * 号，表示注入全部字段，`wire:"-"`  的除外

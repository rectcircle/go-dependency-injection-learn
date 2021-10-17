package main

import "github.com/google/wire"

// Fooer - 接口
type Fooer interface {
	Foo() string
}

// MyFooer - 接口 Fooer 的实现
type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func newMyFooer() *MyFooer {
	foo := MyFooer("Hello, World!")
	return &foo
}

// 结构体 Bar

type Bar string

func newBar(f Fooer) string {
	return f.Foo()
}

var InterfaceSet = wire.NewSet(
	newMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)), // 将结构体类型和接口绑定
	newBar)

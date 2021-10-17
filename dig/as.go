package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"go.uber.org/dig"
)

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

func RunWithInterfaceError1() {
	c := dig.New()

	err := c.Provide(func() (*bytes.Buffer, *bytes.Buffer) {
		return nil, nil
	}, dig.As(new(io.Reader), new(io.Writer)))
	// 错误处理
	if err != nil {
		fmt.Println(c)
		fmt.Printf("错误场景 1 - 构造函数不支持返回多个值: %s\n", err)
	}
}

func RunWithMultipleInterfaceAndName() {
	c := dig.New()

	err := c.Provide(func() *bytes.Buffer {
		return nil
	}, dig.As(new(io.Reader), new(io.Writer)), dig.Name("buffer")) // 通过 as 将结构体和接口进行绑定
	// 错误处理
	if err != nil {
		fmt.Println(c)
		log.Fatal(err)
	}
	fmt.Printf("RunWithMultipleInterfaceAndName: %s", c.String())
}

func RunWithInterfaceMultiple() {
	c := dig.New()

	err := c.Provide(func() (*bytes.Buffer, *bytes.Buffer) {
		return nil, nil
	}, dig.As(new(io.Reader), new(io.Writer))) // 通过 as 将结构体和接口进行绑定
	// 错误处理
	if err != nil {
		fmt.Println(c)
		log.Fatal(err)
	}
	fmt.Printf("返回多个不同接口的值 RunWithInterfaceMultiple: %s", c.String())
}

func RunWithInterfaceError2() {
	c := dig.New()
	// 注册构造函数
	err := c.Provide(newMyFooer, dig.As(new(Fooer)), dig.Group("test")) // 通过 as 将结构体和接口进行绑定
	// 错误处理
	if err != nil {
		fmt.Println(c)
		fmt.Printf("错误场景 2 - `dig.As` 不可以和 `dig.Group` 一起使用: %s\n", err)
	}
}

func RunWithInterfaceError3() {
	c := dig.New()
	// 注册构造函数
	err := c.Provide(newMyFooer, dig.As(new(MyFooer))) // 通过 as 将结构体和接口进行绑定
	// 错误处理
	if err != nil {
		fmt.Println(c)
		fmt.Printf("错误场景 3 - `dig.As` 的参数只允许是接口指针类型: %s\n", err)
	}
}

func RunWithInterface() {

	RunWithInterfaceError1()
	RunWithMultipleInterfaceAndName()
	RunWithInterfaceError2()
	RunWithInterfaceError3()

	c := dig.New()
	// 注册构造函数
	errs := []error{
		c.Provide(newMyFooer, dig.As(new(Fooer))), // 通过 as 将结构体和接口进行绑定
		c.Provide(newBar),
	}
	// 错误处理
	for _, err := range errs {
		if err != nil {
			fmt.Println(c)
			log.Fatal(err)
		}
	}
	// 调用函数，并将 bean 注入函数参数
	err := c.Invoke(func(result string) {
		fmt.Printf("RunWithInterface: %s\n", result)
	})
	if err != nil {
		fmt.Println(c)
		log.Fatalln(err)
	}
	fmt.Println(c.String())
}

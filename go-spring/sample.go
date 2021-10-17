package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-spring/spring-core/gs"
	"github.com/rectcircle/go-dependency-injection-learn/bean/sample"
)

type SampleApp struct {
	C  *sample.C `autowire:""`
	D2 *D2       `autowire:""`
}

type D1 struct {
	D2 *D2 `autowire:""`
}

type D2 struct {
	D1 *D1 `autowire:""`
}

func (a *SampleApp) OnStartApp(e gs.Environment) {
	fmt.Printf("RunSample - gs.AppEvent - OnStartApp: e=%v, c=%s, d2=%v\n", e, a.C, a.D2)
	gs.ShutDown(errors.New(""))
}

func (a *SampleApp) OnStopApp(ctx context.Context) {
	fmt.Printf("RunSample - gs.AppEvent - OnStopApp: %v\n", ctx)
}

func RunSample(a string, b int) {
	// 按照官方的 demo，`gs.Object` 以及 `gs.Provide` 应定义处的 init 函数中

	// go-spring 对 bean 的定义为：一个变量赋值给另一个变量后二者指向相同的内存地址（指针类型）
	// 因此 bean 只有这四种 ptr、interface、chan、func
	gs.Object(&a) // gs.Object(a) // 这种写将报错
	gs.Object(&b) // gs.Object(b) // 这种写法报错

	gs.Provide(func(a *string) *sample.A { return sample.NewA(*a) })
	gs.Provide(func(b *int) *sample.B { return sample.NewB(*b) })
	// 如下两种写法不对，会找不到类型
	// gs.Provide(sample.NewA)
	// gs.Provide(sample.NewB)

	// 循环引用测试
	gs.Object(&D1{})
	gs.Object(&D2{})

	gs.Provide(sample.NewC)

	// 注册 AppEvent
	gs.Provide(new(SampleApp)).Export(new(gs.AppEvent))
	err := gs.Run()
	if err != nil {
		log.Fatal(err)
	}
}

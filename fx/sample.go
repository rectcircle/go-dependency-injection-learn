package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rectcircle/go-dependency-injection-learn/bean/sample"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type SampleIn struct {
	fx.In

	A *sample.A
	B *sample.B
	C *sample.C
}

func RunSample(a string, b int) {

	var sampleC *sample.C
	var sampleIn SampleIn
	var g fx.DotGraph

	app := fx.New(
		// fx.NopLogger, // 关闭日志
		// 配置 fx 框架的日志
		fx.WithLogger(
			func() fxevent.Logger {
				// 默认 log 如下所示
				return &fxevent.ConsoleLogger{W: os.Stderr}
			},
		),
		// fx.ErrorHook(), // 错误处理

		// fx.Supply 等价于 fx.Provide(func() (string, int) { return a, b })
		fx.Supply(a, b),
		// 和 dig 用法类似， dig.Option 能力需通过 fx.Annotated 实现，支持参数对象和结果对象
		// fx.Lifecycle 可以作为构造函数参数
		fx.Provide(sample.NewA, sample.NewB, sample.NewC,
			// 如果想使用命名值和组，可以通过 fx.Annotated 包裹一下
			// 目前还不支持 dig.As 类似的接口绑定
			fx.Annotated{
				Name:   "namedC",
				Target: sample.NewC,
			}),
		// 和 dig 用法类似
		// fx.New 执行完成后，Invoke 就会被调用完成，支持参数对象和结果对象
		// fx.Lifecycle Invoke 函数的参数
		fx.Invoke(func(lc fx.Lifecycle, c *sample.C) {
			fmt.Printf("RunSample - Invoke: %s\n", c)
			// fx.Hook 事件函数，不允许阻塞，默认超时为 fx.DefaultTimeout (15 s)
			// 可以通过 fx.StartTimeout() 和 fx.StopTimeout() 配置
			lc.Append(fx.Hook{
				// 启动回调函数
				OnStart: func(context.Context) error {
					fmt.Printf("RunSample - hooks[0] - OnStart: %s\n", c)
					return nil
				},
				// 停止回调函数
				OnStop: func(context.Context) error {
					fmt.Printf("RunSample - hooks[0] - OnStop: %s\n", c)
					return nil
				},
			})
			// 存在多个，onStart 按照 append 的顺序调用，onSop 按照 append 的逆序调用
			lc.Append(fx.Hook{
				// 启动回调函数
				OnStart: func(context.Context) error {
					fmt.Printf("RunSample -  hooks[1] - OnStart: %s\n", c)
					return nil
				},
				// 停止回调函数
				OnStop: func(context.Context) error {
					fmt.Printf("RunSample -  hooks[1] - OnStop: %s\n", c)
					return nil
				},
			})
		}),
		// 将容器内的通类型的对象赋值给变量，注意，必须是容器内对象的指针类型。也就是说：
		// 如果容器内是 struct 类型，这里传递的是 *struct
		// 如果容器内是 *struct，这里传递的就是**struct
		fx.Populate(&sampleC),
		fx.Populate(&sampleIn), // 不支持 参数对象
		fx.Populate(&g),        // 拿到 DotGraph
	)

	fmt.Printf("RunSample - Populate *sample.C: %s\n", sampleC)
	fmt.Printf("RunSample - Populate sampleIn: %#v\n", sampleIn)
	fmt.Printf("RunSample - DotGraph: \n%s\n", g)

	// app.Run()
	err := app.Start(context.Background())
	fmt.Println(err)
	err = app.Stop(context.Background())
	fmt.Println(err)
}

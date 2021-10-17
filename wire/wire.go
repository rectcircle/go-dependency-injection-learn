//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"io"

	"github.com/google/wire"
	"github.com/rectcircle/go-dependency-injection-learn/bean/sample"
)

// 1. 简单例子
func InitializeSample(aField string, bField int) *sample.C {
	panic(wire.Build(sample.NewA, sample.NewB, sample.NewC))
}

//不能声明在 Injector 里面
var sampleSet = wire.NewSet(sample.NewA, sample.NewB, sample.NewC)

// 2. 使用 sampleSet
func InitializeSample2(aField string, bField int, c bool /*无用输入不会报错*/) *sample.C {
	panic(wire.Build(sampleSet))
}

// 3. 接口和结构体绑定
func InitializeWithInterfaceBind() string {
	panic(wire.Build(InterfaceSet))
}

// 4. 结构体 Provider
func InitializeStructProvider() FooBar2 { // 返回结构体
	panic(wire.Build(StructProviderSet))
}
func InitializeStructProvider2() *FooBar2 { // 返回结构体指针
	panic(wire.Build(StructProviderSet2))
}

// 5. 绑定值
func InitializeValue1() Foo3 {
	panic(wire.Build(BindValueSet1))
}
func InitializeValue2() io.Reader {
	panic(wire.Build(BindValueSet2))
}

// 6. 结构体字段 Provider
func InitializeStructField() string {
	panic(wire.Build(StructFieldProviderSet))
}

// 7. 返回错误
func InitializeTestError(isErr bool) (string, error) {
	panic(wire.Build(newStringOrError))
}

// 8. Cleanup
func InitializeTestCleanup(isErr bool) (string, func(), error) {
	panic(wire.Build(newStringOrCleanup))
}

// 9. 循环依赖，经测试不支持
// func InitializeTestCircularDependencySet(isErr bool) (string, func(), error) {
// 	panic(wire.Build(CircularDependencySet))
// }

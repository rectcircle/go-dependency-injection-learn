package main

import (
	"fmt"
	"log"

	"go.uber.org/dig"
)

type ABEIn struct {
	dig.In // 参数化对象，需嵌入该结构体作为表示，可以用在 Provide 和 Invoke 第一个参数的参数中
	A      int
	B      string
	E      int16 `optional:"true"` // 可选依赖，如果不存在，则为 0 值或者 nil
	// c      bool  `optional:"true"` // dig.In 不允许有私有（未导出）字段，optional 也不行
}

type BCOut struct {
	dig.Out // 结果对象，需嵌入该结构体作为表示，可以用在 Provide 的第一个参数的返回值中使用
	B       string
	C       bool
}

func newBC() BCOut {
	return BCOut{
		B: "b",
		C: true,
	}
}

func newA() int {
	return 1
}

func newD(_ ABEIn) int8 {
	return 2
}

func RunParameterResultObjects() {
	c := dig.New()
	// 注册构造函数
	errs := []error{
		// 调用 newBC 的返回值 BCOut，dig 会将 BCOut 的每个字段作为 value 作为放到容器中，而不是将 BCOut 放入容器中
		c.Provide(newBC),
		c.Provide(newA),
		// dig 会从容器中查询 ABIn 的每个字段，并构建 ABIn 结构体，然后再调用该函数
		c.Provide(newD),
	}
	// 错误处理
	for _, err := range errs {
		if err != nil {
			fmt.Println(c)
			log.Fatal(err)
		}
	}
	// 调用函数，并将 bean 注入函数参数
	err := c.Invoke(func(abe ABEIn, c bool, d int8) {
		// dig 会从容器中查询 ABIn 的每个字段，并构建 ABIn 结构体，然后再调用该函数
		fmt.Printf("RunParameterResultObjects: a=%d, b=%s, c=%t, d=%d, e=%d\n", abe.A, abe.B, c, d, abe.E)
	})
	if err != nil {
		fmt.Println(c)
		log.Fatalln(err)
	}
	fmt.Println(c)
}

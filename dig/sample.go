package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rectcircle/go-dependency-injection-learn/bean/sample"
	"go.uber.org/dig"
)

func RunSample(a string, b int) {
	c := dig.New()
	// 注册构造函数
	errs := []error{
		// 注册构造函数需要的参数
		c.Provide(func() (string, int) { return a, b }),
		// 注册构造函数
		c.Provide(sample.NewA),
		c.Provide(sample.NewB),
		c.Provide(sample.NewC),
	}
	// 错误处理
	for _, err := range errs {
		if err != nil {
			fmt.Println(c)
			log.Fatal(err)
		}
	}
	// 调用函数，并将 bean 注入函数参数
	err := c.Invoke(func(c *sample.C) {
		fmt.Printf("RunSample: %s\n", c)
	})
	if err != nil {
		fmt.Println(c)
		log.Fatalln(err)
	}
	fmt.Println(c.String())
	fmt.Println("print dot graph")
	dig.Visualize(c, os.Stdout)
	fmt.Println()
}

package main

import "fmt"

func main() {
	fmt.Println("=== 简单例子")
	RunSample("string", 1)
	fmt.Println("=== 接口绑定")
	RunWithInterface()
	fmt.Println("=== 参数对象和结果对象")
	RunParameterResultObjects()
	fmt.Println("=== 命名值和组")
	RunNameAndGroup()
}

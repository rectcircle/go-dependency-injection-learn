package main

import (
	"errors"
	"fmt"
)

func newStringOrCleanup(isErr bool) (string, func(), error) {
	if isErr {
		return "", nil, errors.New("模拟构造函数执行抛出异常")
	}
	return "hello World", func() {
		fmt.Println("调用了 cleanup 函数")
	}, nil
}

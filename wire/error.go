package main

import "errors"

func newStringOrError(isErr bool) (string, error) {
	if isErr {
		return "", errors.New("模拟构造函数执行抛出异常")
	}
	return "hello World", nil
}

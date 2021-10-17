package main

import "fmt"

func main() {
	fmt.Printf("call InitializeSample: %s\n", InitializeSample("test", 1))
	fmt.Printf("call InitializeSample2: %s\n", InitializeSample2("test2", 2, true))
	fmt.Printf("call InitializeWithInterfaceBind: %s\n", InitializeWithInterfaceBind())
	fmt.Printf("call InitializeStructProvider: %#v\n", InitializeStructProvider())
	fmt.Printf("call InitializeStructProvider2: %#v\n", InitializeStructProvider2())
	fmt.Printf("call InitializeValue1: %#v\n", InitializeValue1())
	fmt.Printf("call InitializeValue2: %#v\n", InitializeValue2())
	fmt.Printf("call InitializeStructField: %s\n", InitializeStructField())
	s, e := InitializeTestError(true)
	fmt.Printf("call InitializeTestError(true): %s, %v\n", s, e)
	s, e = InitializeTestError(false)
	fmt.Printf("call InitializeTestError(false): %s, %v\n", s, e)
	s, cleanup, e := InitializeTestCleanup(true)
	fmt.Printf("call InitializeTestCleanup(true): %s, %v\n", s, e)
	if cleanup != nil {
		cleanup()
	}
	s, cleanup, e = InitializeTestCleanup(false)
	fmt.Printf("call InitializeTestCleanup(false): %s, %v\n", s, e)
	if cleanup != nil {
		cleanup()
	}
}

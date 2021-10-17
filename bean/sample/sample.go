package sample

import "fmt"

// 类型 A 和 构造器

type A struct {
	aField string
}

func NewA(aField string) *A {
	return &A{
		aField: aField,
	}
}

// 类型 B 和 构造器

type B struct {
	bField int
}

func NewB(bField int) *B {
	return &B{
		bField: bField,
	}
}

// 类型 B 和 构造器

type C struct {
	a *A
	b *B
}

func (c *C) String() string {
	if c == nil {
		return "sample.C<nil>"
	}
	return fmt.Sprintf("I am C and c.a is %s, c.b is %d", c.a.aField, c.b.bField)
}

func NewC(a *A, b *B) *C {
	return &C{
		a: a,
		b: b,
	}
}

// 假设最终要构造一个 C，写法如下

func ManualInitialize(aField string, bField int) *C {
	a := NewA(aField)
	b := NewB(bField)
	c := NewC(a, b)
	return c
}

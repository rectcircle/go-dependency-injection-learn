package main

import (
	"fmt"
	"log"

	"go.uber.org/dig"
)

func newInt() int {
	return -1
}

func newInt0And1() (int, int) {
	return 0, 1
}

func newIntAndString() (int, string) {
	return 1, "a"
}

func newInt2() int {
	return 2
}

func newInt3() int {
	return 3
}

func newItemA() string {
	return "a"
}

type Int3AndItemBOut struct {
	dig.Out
	Int4      int      `name:"int4"`
	ItemB     string   `group:"list"`
	ItemOther []string `group:"list,flatten"` // 注意展平
}

func newItemCD() (string, string) {
	return "c", "d"
}

func newInt4AndItemBOut() Int3AndItemBOut {
	return Int3AndItemBOut{
		Int4:      4,
		ItemB:     "b",
		ItemOther: []string{"e", "f", "g"},
	}
}

type NameAndGroupIn struct {
	dig.In
	Int0 int      `name:"int0" optional:"true"`
	Int1 int      `name:"int1" optional:"true"`
	Int2 int      `name:"int2"`
	Int3 int      `name:"int3"`
	Int4 int      `name:"int4"`
	List []string `group:"list"`
}

type IntStringGroup struct {
	dig.In
	ListInt        []int    `group:"list_int"`
	ListIntString1 []int    `group:"list_int_string"`
	ListIntString2 []string `group:"list_int_string"`
	ListString     []string `group:"list_string"` // 允许不存在
}

func RunNameAndGroupError1() {
	c := dig.New()
	// 注册构造函数
	errs := []error{
		c.Provide(newInt0And1),
		c.Provide(newInt0And1, dig.Name("int0")),
		c.Provide(newInt0And1, dig.Name("int0"), dig.Name("int0")),
	}
	// 错误处理
	for i, err := range errs {
		if err != nil {
			fmt.Println(c)
			fmt.Printf("错误场景1[%d] - 不支持在同一个函数里返回同一类型的多个值: %s\n", i, err)
		}
	}
}

func RunNameAndGroupError2() {
	c := dig.New()
	// 注册构造函数
	errs := []error{
		c.Provide(newInt2, dig.Name("int2")),
	}
	// 错误处理
	for _, err := range errs {
		if err != nil {
			fmt.Println(c)
			log.Fatal(err)
		}
	}
	err := c.Invoke(func(int_ int) {
		fmt.Printf("int: %d\n", int_)
	})
	if err != nil {
		fmt.Println(c)
		fmt.Println("错误场景2 - 命名值不能注入到非命令名参数中:", err)
	}
}

func RunNameAndGroupError3() {
	c := dig.New()
	// 注册构造函数
	errs := []error{
		c.Provide(newInt2, dig.Name("int2")),
		c.Provide(newInt2, dig.Name("int2")),
	}
	// 错误处理
	for _, err := range errs {
		if err != nil {
			fmt.Println(c)
			fmt.Printf("错误场景3 - 不允许同一类型存在多个相同的命名 %s\n", err)
		}
	}
}

func RunNameReturnMultipleAndNameMultiple() {
	c := dig.New()
	err := c.Provide(newIntAndString, dig.Name("a"), dig.Name("b"))
	if err != nil {
		fmt.Println(c)
		log.Fatal(err)
	}
	fmt.Println(c)
	fmt.Println("RunNameReturnMultipleAndNameMultiple: int 和 string 每个值都会被命名成 b")
}

func RunGroupReturnTypeGroup() {
	c := dig.New()

	// 注册构造函数
	errs := []error{
		c.Provide(newInt0And1, dig.Group("list_int")),
		c.Provide(newIntAndString, dig.Group("list_int_string")),
	}
	// 错误处理
	for _, err := range errs {
		if err != nil {
			fmt.Println(c)
			log.Fatal(err)
		}
	}
	err := c.Invoke(func(IntGroup IntStringGroup) {
		fmt.Printf("RunGroupReturnSameTypeGroup: list_int=%#v\n", IntGroup)
	})
	if err != nil {
		fmt.Println(c)
		log.Fatalln(err)
	}
	fmt.Println(c)
	fmt.Println("RunGroupReturnSameTypeGroup: 构造函数返回想同类型可以使用 group 进行聚合")
}

func RunNameAndGroup() {

	RunNameAndGroupError1()
	RunNameAndGroupError2()
	RunNameAndGroupError3()
	RunNameReturnMultipleAndNameMultiple()
	RunGroupReturnTypeGroup()

	c := dig.New()

	// 注册构造函数
	errs := []error{
		c.Provide(newInt),
		c.Provide(newInt2, dig.Name("int2")),
		c.Provide(newInt3, dig.Name("int3")),
		c.Provide(newInt4AndItemBOut),
		c.Provide(newItemA, dig.Group("list")),
		c.Provide(newItemCD, dig.Group("list")),
	}
	// 错误处理
	for _, err := range errs {
		if err != nil {
			fmt.Println(c)
			log.Fatal(err)
		}
	}
	// 调用函数，并将 bean 注入函数参数
	err := c.Invoke(func(_int int, in NameAndGroupIn) {
		fmt.Printf("RunNameAndGroup: int=%d, in=%#v\n", _int, in)
	})
	if err != nil {
		fmt.Println(c)
		log.Fatalln(err)
	}
	fmt.Println(c.String())
}

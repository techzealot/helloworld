package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

// 包内部变量必须使用var定义，无全局变量
var aa = 1

// () 可以省略多个var
var (
	bb = 2
	cc = 3
	dd = 4
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func variableInitValue() {
	var a, b int = 3, 4
	ch := `sdd
	dsdsd
sdsds
\r\n`
	var s string = "abc"
	fmt.Println(a, b, s, ch)
}

func variableTypeDeduction() {
	a, b, c, s := 3, 4, "abc", true
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// := 定义并赋值,推荐使用 代码越少越简单越好
	a, b, c, s := 3, 4, "abc", true
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i

	fmt.Println(cmplx.Abs(c))

	eu := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(eu)

	eu = cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%3f \n", eu)
}

// 类型转换必须显式声明
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

const (
	d = 1
	e = 2
)

func consts() {
	const filename = "abc.txt"
	// 常量可以当做各种类型使用，相当于文本替换
	const a, b = 3, 4
	var c int
	// 此处无需强转
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	/*const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)*/
	const (
		cpp    = iota // 0
		_             // 1
		python        // 2
		golang        // 3
	)
	// b kb mb gb tb pb
	const (
		b  = 1 << (10 * iota) // 1 << (10*0)
		kb                    // 1 << (10*1)
		mb                    // 1 << (10*2)
		gb                    // 1 << (10*3)
		tb                    // 1 << (10*4)
		pb                    // 1 << (10*5)
	)
	fmt.Println(b, kb, mb, gb, tb, pb) // 1 1024 1048576 1073741824 1099511627776 1125899906842624

	fmt.Println(cpp, python, golang) // 0 2 3
}

func main() {
	fmt.Println(runtime.GOARCH)
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc, dd)
	euler()
	triangle()
	enums()
}

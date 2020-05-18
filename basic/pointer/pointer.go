package main

import "fmt"

func pa() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a, pa, *pa, &pa, &a) //3 0xc00000a0e8 3 0xc000006028 0xc00000a0e8
}

func passByValue(a int) {
	a++
}
func passByRef(a *int) {
	*a++
}

//推荐做法
func swapNoSideEffect(a, b int) (int, int) {
	return b, a
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	a := 2
	passByValue(a)
	fmt.Println(a) //2
	passByRef(&a)
	fmt.Println(a) //3

	c, d := 5, 6
	swap(&c, &d)
	fmt.Println(c, d) //6,5

	c, d = swapNoSideEffect(c, d)
	fmt.Println(c, d) //5,6
	var e int
	c, d, e = d, c, d/c
	fmt.Println(c, d, e) //6 5 1 将右侧变量先运算替换，然后一次赋多值
}

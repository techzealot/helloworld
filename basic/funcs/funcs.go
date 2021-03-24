package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 不要滥用多值返回，一般最后一个为error
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func evalSilent(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "%":
		_, r := div(a, b)
		return r, nil
	default:
		return 0, fmt.Errorf("unsupport op : %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d,%d) \n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(div(13, 3))
	if result, err := evalSilent(13, 3, "x"); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(result)
	}
	// 匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

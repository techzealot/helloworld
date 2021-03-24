package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score %d \n", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 0o0:
		g = "A"
	}
	return g
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported op:" + op)
	}
}

func loop() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
}

func main() {
	const filename = "abc.txt"
	// file作用域在if语句内
	if file, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", file)
	}
}

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		fmt.Println("sum=", sum)
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		if i == 1 {
			fmt.Println(a(i))
		} else {
			fmt.Println(a(i))
		}

	}
}

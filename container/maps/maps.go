package main

import "fmt"

func deleteMap(m map[string]string) {
	delete(m, "name")
}

func main() {
	// 1
	m := map[string]string{
		"name": "tom",
		"age":  "20",
	}
	fmt.Println(m) // map[age:20 name:tom]

	// 2 或m2 := make(map[string]int,cap)
	m2 := make(map[string]int)
	fmt.Println(m2) // map[]

	m2["a"] = 1
	fmt.Println(m2)
	// 3
	var m3 map[string]int
	fmt.Println(m3) // map[]

	//** this will panic
	m3["a"] = 1

	//** this will not panic now
	delete(m3, "test")

	for k := range m {
		fmt.Println(k)
	}
	s := "test"
	if v, ok := m[s]; ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println(v, ok)
	}
	deleteMap(m)
	fmt.Println("after delete:", m)
}

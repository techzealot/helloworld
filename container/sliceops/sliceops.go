package main

import "fmt"

func PrintSlice(s []int) {
	fmt.Printf("%v ,len= %d ,cap= %d \n", s, len(s), cap(s))
}

func main() {
	var s []int // zero value for slice is []

	PrintSlice(s)

	s1 := []int{2, 4, 6, 8}

	s2 := make([]int, 16)

	s3 := make([]int, 10, 32)

	PrintSlice(s2)

	PrintSlice(s3)

	copy(s2, s1)

	fmt.Println(s2) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]

	// 删除下标为3的元素
	s2 = append(s2[:3], s2[4:]...)

	PrintSlice(s2) //[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0] ,len= 15 ,cap= 16
}

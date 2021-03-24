package main

import "fmt"

// 将会更改slice对应的数组元素
func updateSlice(s []int) {
	// 会更改slice的第0个元素对应的数组元素
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	s2 := s1[3:5]

	// arr[8]被覆盖为10
	s3 := append(s2, 10)

	s4 := append(s3, 10)

	s5 := append(s4, 10)

	// 对切片的插入一旦超过底层数组，便会生成一个更大（2len(arr)）的新数组并会拷贝原来数组的数据
	// s4 s5 no longer view arr
	fmt.Println("s3 ,s4 ,s5 =", s3, s4, s5) // s3 ,s4 ,s5 = [5 6 10] [5 6 10 10] [5 6 10 10 10]

	fmt.Println(arr) //[0 1 2 3 4 5 6 10]
}

package main

import "fmt"

// 数组的值会拷贝过来，原数组不会被改变
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayWithEffect(arr *[5]int) {
	// 此处等价于 (*arr)[0] = 100,语法处理，arr并不是数组的头指针（不同于C）
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 默认初始化
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	// 编译器推断
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3) //[0 0 0 0 0] [1 3 5] [2 4 6 8 10]
	// 4行5列二维数组
	var grade [4][5]bool
	fmt.Println(grade) //[[false false false false false] [false false false false false] [false false false false false] [false false false false false]]

	// 强烈不推荐，其他语言的做法
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	// 推荐做法,可同时获取索引和值,_可以省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}
	printArray(arr3)
	fmt.Println(arr3) //[2 4 6 8 10]
	printArrayWithEffect(&arr3)
	fmt.Println(arr3) //[100 4 6 8 10]
}

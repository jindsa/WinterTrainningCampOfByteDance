package main

import "fmt"

func main() {

	var a [5]int
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c [10]int
	for i := 0; i < 10; i++ {
		c[i] = i + 1
		fmt.Println(c[i])
	}
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/*
数组：具有编号且长度固定的元素序列。
更多用的是切片
*/

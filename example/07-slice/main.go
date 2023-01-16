package main

import "fmt"

func main() {

	// 初始化一个字符串切片
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3
	// append 给切片添加元素
	/*
		切片原理是：它存储了一个长度和一个容量，加一个指向数组
		指针，在执行append操作时，若容量不够或扩容并且返回新的
		slice。
		类似python的切片，比如代表取出第二个到第五个位置的元素
		时不包括第五个元素，但不支持负数索引。
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
	sports := make([]string, 3)
	sports = append(sports, "basketball")
	sports[1] = "football"
	sports[0] = "pingpong"
	sports[2] = "swimming"
	fmt.Println(sports, len(sports))
}

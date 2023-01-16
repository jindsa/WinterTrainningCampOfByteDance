package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3
	var res2 = float32(add2(3, 6))
	fmt.Println(res2)
	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}

/*
golang的变量类型是后置的；
golang的函数原生支持返回多个值。实际逻辑代码中几乎所有函数
都返回两个值：真正的返回结果；错误信息

*/

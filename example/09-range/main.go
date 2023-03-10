package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", m[k]) // key a; key b
	}
}

/*
range 可用来快速遍历slice或map。
range遍历时，对于数组会返回两个值：索引和对应位置的值（不需要索引可以用下划线忽略）

*/

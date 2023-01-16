package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 10
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknow"]) // 0

	r, ok := m["unknow"]
	fmt.Println(r, ok) // 0 false
	k, v := m["two"]
	var a, b = m["one"]
	fmt.Println(a, b, "\\n")
	fmt.Println(k, v)

	delete(m, "one")
	fmt.Println(m)
	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}

/*
map 叫做哈希或字典
用make(map[string]int)创建一个字典:key是string类型；value是int类型
可用delete删除键值对

golang的map完全无序，遍历时是随机顺序的。

*/

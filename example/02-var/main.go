package main

import (
	"fmt"
	"math"
)

func main() {

	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "jinzhiling"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}

// Go的字符串是内置的类型：“+”拼接，“=”比较两个字符串
/*
变量的声明：
1.通过  var name string =""  一般自动推导变量类型；
2.形式： 变量 :=值  如 name:="jzl"

常量的声明：
就是将var改成const。
go语言中的常量没有确定的类型，自动根据使用的上下文确定类型。



*/

package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

/*
与C语言比较：
1.switch后面的变量不需要小括号；
2.在C语言中，switch case如果不加break的话，会继续向下跑完所有的case
在go语言中，不需要添加任何的break的。
3.go的switch功能更强大。可以使用任意的变量类型，甚至可以用来
取代任意的if else语句。如：可以在switch后不加任何的变量，在case后面
写条件分支。
*/

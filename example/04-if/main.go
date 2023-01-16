package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	if grade := 99; grade == 100 {
		fmt.Println(grade, "is not perfect.")
	} else if grade >= 90 {
		fmt.Println(grade, "is  rather good.")
	} else if grade >= 60 {
		fmt.Println(grade, "is just soso.")
	} else {
		fmt.Println("you are a little chicken.")
	}
}

/*
if-else:
与C语言相比：
1.if中的条件不需要小括号；条件之后必须有大括号
if a>10 {  }
2.在if条件中,条件中如果定义变量，定义之后的条件用 “;” 分隔。
if num:=89;num>90{

}else if num>80{

}else{

}

*/

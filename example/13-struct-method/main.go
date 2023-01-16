package main

import "fmt"

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
}

/*
结构体函数优点类似类成员函数
可以这样调用结构体函数：
a.checkPassword("xx")
具体的结构体函数定义：将第一个参数加上括号写到函数名称的前面
实现结构体方法有两种写法：带指针、不带指针
如果带指针那就是对这个结构体进行修改；不带指针实际上操作的
是一个结构体的拷贝。
*/

package main

import "fmt"

type user struct {
	name     string
	password string
	sex      string
}

func main() {
	a := user{name: "wang", password: "1024", sex: "unknown"}
	b := user{"wang", "1024", "male"}
	c := user{name: "wang"}
	c.password = "1024"
	c.sex = "female"
	var d user
	d.name = "wang"
	d.password = "1024"
	d.sex = "gay"
	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "1024"))   // false
	fmt.Println(checkPassword2(&a, "1024")) // false
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}

/*
结构体指针可以大大节省大的结构体拷贝开销
*/

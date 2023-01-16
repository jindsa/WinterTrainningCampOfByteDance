package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string `json:"xingming"`
	Age   int    `json:"age"`
	Hobby []string
}

func main() {
	a := userInfo{Name: "wan", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)         // [123 34 78 97...]
	fmt.Println(string(buf)) // {"Name":"wang","age":18,"Hobby":["Golang","TypeScript"]}

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}

/*
对于一个已有的结构体，只要保证每个字段的第一个字母大写，也就是
公开字段。哪个这个结构体就可以用json.marshaler去序列化变成
一个json的字符串。

序列化之后的字符串也可以用json.unmarshaler反序列化到一个空的变量

默认序列化之后的字符串大写字母开头而非下划线；可以在结构体的成员
变量后面添加json tag等语法去修改json结果里面的字段名。

*/

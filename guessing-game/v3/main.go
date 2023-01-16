package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)

	/*
	    bufio.NewReader将一个文件转化为reader变量；
	   reader变量上有许多用来操作一个流的操作：
	   1.它的ReaderString方法来读取一行：如果读取失败将打印错误信息
	   并退出读入操作；若读取成功，该方法的返回结果包含结尾的换行符，
	   我们可以将其去点，再转换成数字，如果转换失败同样打印错误信息并
	   退出。
	   2.
	*/
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	input = strings.Trim(input, "\r\n")

	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println("You guess is", guess)
}

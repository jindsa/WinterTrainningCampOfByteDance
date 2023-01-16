package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100)
	fmt.Println("The secret number is ", secretNumber)
}

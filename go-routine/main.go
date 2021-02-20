package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world")
}

func main() {
	go hello()
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
}

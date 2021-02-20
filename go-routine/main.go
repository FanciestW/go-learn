package main

import (
	"fmt"
)

// A Go routine function to say hello world
func hello(done chan bool) {
	fmt.Println("Hello world")
	done <- true
}

func main() {
	var done chan bool = make(chan bool)
	go hello(done)
	fmt.Println("main function")
	<-done
}

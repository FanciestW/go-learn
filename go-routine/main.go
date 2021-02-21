package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	one = iota
)

// A Go routine function to say hello world
func hello(done chan bool) {
	fmt.Println("Hello world")
	var nameChan chan string = make(chan string)
	go getName(nameChan)
	receivedName := <-nameChan
	fmt.Printf("Welcome, %s\n", receivedName)
	done <- true
}

func getName(name chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your name: ")
	text, _ := reader.ReadString('\n')
	name <- text
}

func main() {
	var done chan bool = make(chan bool)
	go hello(done)
	fmt.Println("main function")
	<-done
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	nameReader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name: ")
	name, _ := nameReader.ReadString('\n')
	// convert CRLF to LF
	name = strings.Replace(name, "\n", "", -1)

	var message string
	var error error
	message, error = greetings.Hello(name)

	// If an error was returned, print it to the console and
	// exit the program.
	if error != nil {
		log.Fatal(error)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println(quote.Go())
}

package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var message map[string]string
	var error error
	message, error = greetings.Hellos([]string{"William", "Alyson"})

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

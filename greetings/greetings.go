package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting string for a person's name. (Capitalized function name is Go exported name function)
func Hello(name string) (string, error) {
	// If no name given, return empty string with an error
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats. (No size is specified so it is a slice)
	formats := []string{"Hi %v, welcome!", "Great to see you, %v!", "Hello there, %v!"}

	// Return a randomly selected message format by specifying a random
	// index of the slice of formats.
	return formats[rand.Intn(len(formats))]
}

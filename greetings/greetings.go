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

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	var messages map[string]string
	messages = make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, errors.New("An empty name was received")
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
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

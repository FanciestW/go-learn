package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting string for a person's name. (Capitalized function name is Go exported name function)
func Hello(name string) (string, error) {
	// If no name given, return empty string with an error
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

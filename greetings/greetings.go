package greetings

import "fmt"

// Hello returns a greeting string for a person's name. (Capitalized function name is Go exported name function)
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

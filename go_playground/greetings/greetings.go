package greetings

import "fmt"

func Greeting(name string) string {
	message := fmt.Sprintf("hi %v. Welcome!", name)
	return message
}
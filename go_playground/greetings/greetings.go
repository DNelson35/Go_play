package greetings

import ( 
	"errors"
	"fmt"
)
func Greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("name must not be empty")
	}
	message := fmt.Sprintf("hi %v. Welcome!", name)
	return message, nil
}
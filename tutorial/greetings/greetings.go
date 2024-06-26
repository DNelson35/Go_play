package greetings

import ( 
	"errors"
	"fmt"
	"math/rand"
)
func Greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("name must not be empty")
	}
	message := fmt.Sprintf(randFormat(), name)
	return message, nil
}


func MultiGreet(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greeting(name)
		
		if err != nil {
			return nil, err
		}
		messages[name] = message

	}

	return messages, nil
}



func randFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
		"%v, who is you?",
	}

	return formats[rand.Intn(len(formats))]
}
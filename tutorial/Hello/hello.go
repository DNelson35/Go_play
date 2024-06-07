package main

// import "fmt"

// import "rsc.io/quote"

import (
	"fmt"
	"log"

	"example.com/greetings"
)


var name string = "Damien"

func main() {
	// fmt.Printf("Hello, World! My name is '%s'\n", name)
	// fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Damien", "Felesia", "Juniper", "Atlas"}

	messages, err := greetings.MultiGreet(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages{
		fmt.Println(message)
	}
	
}

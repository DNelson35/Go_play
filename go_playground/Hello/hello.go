package main

// import "fmt"

// import "rsc.io/quote"

import (
	"fmt"
	"log"

	"example.com/greetings"
)


// var name string = "Damien"

func main() {
	// fmt.Printf("Hello, World! My name is '%s'\n", name)
	// fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Greeting("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

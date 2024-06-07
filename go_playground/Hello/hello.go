package main

// import "fmt"

// import "rsc.io/quote"

import (
	"fmt"

	"example.com/greetings"
)


var name string = "Damien"
func main() {
	// fmt.Printf("Hello, World! My name is '%s'\n", name)
	// fmt.Println(quote.Go())
	message := greetings.Greeting(name)

	fmt.Println(message)
}

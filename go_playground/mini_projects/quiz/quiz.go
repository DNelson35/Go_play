package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Printf("Would you like to play (y/n)? ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input) // Remove trailing newline

	if input == "n" || input == "" {
		fmt.Println("\nSorry to hear that. Maybe next time, goodbye.")
		return
	}

	fmt.Println("Awesome, let's get started...")

	time.Sleep(2 * time.Second)

	clearScreen()

	fmt.Printf("Question 1: Who am I? ")
	input, err = reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error reading input:", err)
		return
	}
	input = strings.TrimSpace(input) // Remove trailing newline

	if input != "Damien" {
		fmt.Printf("Sorry, %s is incorrect.\n", input)
		fmt.Printf("Try again next time...\n")
		time.Sleep(2 * time.Second)
		return
	}

	fmt.Println("You answered correctly.")
	// Add logic here to continue with the next question or exit the game
}

func clearScreen() {
	fmt.Print("\033c") // ANSI escape sequence to clear the screen
}

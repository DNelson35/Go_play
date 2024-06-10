package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to my quiz game!") // Println adds newline at the end
	fmt.Printf("Would you like to play (y/n)? ") // printf allows variables and continues in same line
	reader := bufio.NewReader(os.Stdin) // look up how the buffer works and what the reder is and how and why this is better than simply scanning in text form stdin
	input, err := reader.ReadString('\n') // telling the reader to read stdin up untill new line
	if err!= nil {
		fmt.Println("Error reading input:", err) // if error print message quit program
		return
	}
	input = strings.TrimSpace(input) // Remove trailing newline

	if strings.ToLower(input) != "y" { // exit program if they don't answer yes
		fmt.Println("\nSorry to hear that. Maybe next time, goodbye.")
		return
	}

	questions := map[string]string{
		"how many moons does Earth have? ": "1",
		"how long is a day in hours? ": "24",
		"who is the flash? ": "barry alan",
		"what year was the declaration of independence signed? ": "1776",
	}

	fmt.Println("Awesome, let's get started...")

	time.Sleep(2 * time.Second) // if yes we print message wait 2 seconds then clear screen

	clearScreen()
    score := 0

	for question, answer := range questions{
		fmt.Printf("Score = %v\n", score)

		fmt.Printf("%v", question)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			time.Sleep(time.Second * 2)
			return
		}

		input = strings.TrimSpace(input)
		

		if strings.ToLower(input) != answer {
			fmt.Printf("Sorry, %s is incorrect.\n", input)
			fmt.Printf("Try again next time...\n")
			time.Sleep(2 * time.Second)
			clearScreen()
		} else  {
			fmt.Println("You answered correctly.")
			score ++
			time.Sleep(2 * time.Second)
			clearScreen()
		}
		
	}
    fmt.Printf("Final Score = %v out of %v\n", score, len(questions))
}

func clearScreen() {
	fmt.Print("\033c") // ANSI escape sequence to clear the screen
}


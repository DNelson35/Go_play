package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Printf("would you like to play (y/n)? ")
	var input string
	fmt.Scan(&input)

	if input == "n" || input == ""{
		fmt.Println("\nsorry to hear that maybe next time, goodbye.")
		return
	}

	fmt.Println("Awesome lets get started...")

	time.Sleep(2 * time.Second)

	fmt.Println("\033[2J")

	fmt.Printf("question 1: who am I? ")
	fmt.Scan(&input)

	if input != "damien"{
		fmt.Printf("sorry %v is incorrect\n", input)
		fmt.Printf("try again next time...")

		time.Sleep(2 * time.Second)
		return
	}





}
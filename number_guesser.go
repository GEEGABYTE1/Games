package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	computer_number := rand.Intn(10)

	var user_input int
	fmt.Println("Please type in a number to guess what the computer chose from 1-10: ")
	fmt.Scan(&user_input)

	if user_input == computer_number {
		fmt.Printf("Congratulations, your number was %v and the computer's number was also %v", user_input, computer_number)
		fmt.Println(" ")
	} else if user_input > computer_number {
		fmt.Printf("Oops, looks like your number, %v, and was greater than the computer's number, %v", user_input, computer_number)
	} else {
		fmt.Printf("Oops, looks like your number, %v, and was smaller than the computer's number, %v", user_input, computer_number)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	rock := 1
	paper := 2
	scissors := 3

	fmt.Printf("Type %d for rock ", rock)
	fmt.Println(" ")
	fmt.Printf("Type %d for paper ", paper)
	fmt.Println(" ")
	fmt.Printf("Type %d for scissors ", scissors)
	fmt.Println(" ")

	computer_input := rand.Intn(3)

	var user_input int
	fmt.Scan(&user_input)

	if user_input == computer_input {
		fmt.Println("You have tied with the computer!")
	} else if user_input == 1 && computer_input == 3 {
		fmt.Println("You have chosen rock and the computer chose scissors!")
		fmt.Println("Congratulations, you have won")
	} else if user_input == 3 && computer_input == 1 {
		fmt.Println("You have chosen scissors and the computer chose rock!")
		fmt.Println("Sadly, you have lost")
	} else if user_input == 1 && computer_input == 2 {
		fmt.Println("You have chosen rock and the computer chose paper!")
		fmt.Println("Sadly, you have lost")
	} else if user_input == 2 && computer_input == 1 {
		fmt.Println("You have chosen paper and the computer chose rock!")
		fmt.Println("Congratulations, you have won")
	} else if user_input == 2 && computer_input == 3 {
		fmt.Println("You have chosen paper and the computer chose scissors!")
		fmt.Println("Sadly, you have lost")
	} else if user_input == 3 && computer_input == 2 {
		fmt.Println("You have chosen scissors and the computer chose paper!")
		fmt.Println("Congratulations, you have won")
	}
}

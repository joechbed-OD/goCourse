package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generating a random number between 1 and 100
	target := random.Intn(100) + 1

	//displaying a welcome message
	fmt.Println("WELCOME TO THE GUESSING GAME!")
	fmt.Println("I have chosen a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		//checking if the quess is correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the number!")
			break
		} else if guess < target {
			fmt.Println("Too low!, Try guessing a higher number.")
		} else{
			fmt.Println("Too high!, Try guessing a lower number.")
		}
	}
}
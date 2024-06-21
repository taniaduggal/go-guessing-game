package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//create a secret number
	secret := getRandomNumber()
	// fmt.Println(secret)

	for matching := false; !matching; {
		guess := getUserInput()
		fmt.Println("the secret is:", secret, "the guess is:", guess)

		//make comparison (Secret vs guess)
		matching = isMatching(secret, guess)
		// fmt.Println(matching)
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your guess is too big")
		return false
	} else if guess < secret {
		fmt.Println("Your guess is too small")
		return false
	} else {
		fmt.Println("You win!")
		return true
	}
}

func getUserInput() int {
	fmt.Print("Please input your guess:")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed:", input)
	}
	return input
}

func getRandomNumber() int {
	return rand.Int() % 11
}

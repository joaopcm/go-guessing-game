package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println("Pick a number between 1 and 100.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("Enter your guess:")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			return
		}

		switch {
			case guessInt < x:
				fmt.Println("You guessed too low.")
			case guessInt > x:
				fmt.Println("You guessed too high.")
			case guessInt == x:
				fmt.Println("You guessed correctly! The number was", x)
				return
		}

		guesses[i] = guessInt
	}

	fmt.Printf("You guessed %d times. The number was %d. The guesses were %v\n", len(guesses), x, guesses)
}
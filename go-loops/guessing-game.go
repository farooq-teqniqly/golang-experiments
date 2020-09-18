package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	var guessCount = 1

	for {
		if guessCount == 1 {
			fmt.Print("Guess an number between 1 and 100: ")
		} else {
			fmt.Printf("Guess a number (%v guesses remaining): ", 10-guessCount)
		}

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Thank you for playing!")
			break
		}

		guess, err := strconv.Atoi(input)
		guessCount++

		if guess == target {
			fmt.Println("You guessed correctly!")
			fmt.Println("Thank you for playing!")
			break
		} else {
			if guess > target {
				fmt.Println("Your guess is too high.")
			} else {
				fmt.Println("Your guess is too low.")
			}

			if guessCount == 10 {
				fmt.Printf("You ran out of guesses. The number was %v.\n", target)
				fmt.Println("Thank you for playing!")
				break
			}
		}
	}
}

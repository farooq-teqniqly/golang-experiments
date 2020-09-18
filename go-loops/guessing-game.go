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
	var guess_count int

	for {
		if guess_count == 0 {
			fmt.Print("Guess an number between 1 and 100: ")
		} else {
			fmt.Print("Guess a number (" + fmt.Sprint(10-guess_count) + " guesses remaining).")
		}

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Good bye.")
		}

		guess, err := strconv.Atoi(input)
		guess_count++

		if guess == target {
			fmt.Println("You guessed correctly!")
		} else {
			fmt.Println("You guessed incorrectly.")

			if guess_count == 10 {
				fmt.Println("Game over!")
			}
		}
	}
}

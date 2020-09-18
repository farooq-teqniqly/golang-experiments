package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var err error
	var input string

	for {
		fmt.Print("Enter a grade: ")

		reader := bufio.NewReader(os.Stdin)

		input, err = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		grade, err := strconv.ParseFloat(input, 64)

		if err != nil {
			log.Fatal(err)
		}

		var status string

		if grade >= 60 {
			status = "passing"
		} else {
			status = "failing"
		}

		fmt.Println(status)
	}

	fmt.Println("Good bye.")
}

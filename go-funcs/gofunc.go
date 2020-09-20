package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the wall width: ")
	widthStr, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
		return
	}

	width, err := strconv.ParseFloat(strings.TrimSpace(widthStr), 64)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Print("Enter the wall height: ")
	heightStr, err := reader.ReadString('\n')

	height, err := strconv.ParseFloat(strings.TrimSpace(heightStr), 64)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Print("How many paint buckets do you have? ")
	bucketsStr, err := reader.ReadString('\n')

	buckets, err := strconv.Atoi(strings.TrimSpace(bucketsStr))

	if err != nil {
		log.Fatal(err)
		return
	}

	paintNeeded, err := paintNeeded(height, width, buckets)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%0.2f liters needed", paintNeeded)
}

func calcArea(x float64, y float64) float64 {
	return x * y
}

func paintNeeded(x float64, y float64, buckets int) (float64, error) {
	if x < 0 || y < 0 {
		err := errors.New("the height and width must be positive numbers")
		return 0, err
	}

	return calcArea(x, y) / float64(buckets), nil
}

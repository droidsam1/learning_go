package main

import (
	"fmt"
	"math/rand"
)

func main() {

	oneHundredNumbers := create()

	for _, i := range oneHundredNumbers {
		switch {
		case i == 0:
			fmt.Println("Nevermind")
		case i%2 == 0 && i%3 == 0:
			fmt.Println("Six")
		case i%2 == 0:
			fmt.Println("Two")
		case i%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Nevermind")
		}
	}

}

func create() []int {

	oneHundredNumbers := []int{}

	for i := 0; i < 100; i++ {
		oneHundredNumbers = append(oneHundredNumbers, rand.Intn(100))
	}

	return oneHundredNumbers
}

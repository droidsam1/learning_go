package main

import "fmt"

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

	for i := range 100 {
		oneHundredNumbers = append(oneHundredNumbers, i)
	}

	return oneHundredNumbers
}

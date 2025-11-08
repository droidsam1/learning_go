package main

import (
	"fmt"
	"math/rand"
)

func main() {

	oneHundredNumbers := []int{}

	for i := 0; i < 100; i++ {
		oneHundredNumbers = append(oneHundredNumbers, rand.Intn(100))
	}

	fmt.Println(oneHundredNumbers)
}

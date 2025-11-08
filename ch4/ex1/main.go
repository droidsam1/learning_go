package main

import "fmt"

func main() {

	oneHundredNumbers := []int{}

	for i := 0; i < 100; i++ {
		oneHundredNumbers = append(oneHundredNumbers, i)
	}

	fmt.Println(oneHundredNumbers)

}

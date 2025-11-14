package main

import "fmt"

func main() {
	//creates a slice of ints using make([]int, 0, 5)
	numbers := make([]int, 0, 5)
	numbers = append(numbers, 10, 20, 30)

	fmt.Println("length: ", len(numbers), "capacity: ", cap(numbers), "content: ", numbers)

	s2 := numbers[:2]
	s2 = append(s2, 999)

	fmt.Println("s:", numbers, "s2:", s2)

	// s2 and s shared the same underleying array. Since s2 does not exceed the original capacity, s2 is reusing the same underlying array and therefore when using append it is replacing the last element of the original array
}

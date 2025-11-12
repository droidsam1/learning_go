package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	// the problem is that total is being redeclared in the for loop block shadowing the variable from the outer block
}

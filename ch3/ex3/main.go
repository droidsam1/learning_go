package main

import "fmt"

func main() {

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	founder := Employee{"john", "doe", 1}
	coFounder := Employee{
		firstName: "jane",
		lastName:  "doe",
		id:        2,
	}
	var firstWorker Employee

	firstWorker.firstName = "someDude"
	firstWorker.lastName = "withSomeLastName"
	firstWorker.id = 3

	fmt.Println(founder)
	fmt.Println(coFounder)
	fmt.Println(firstWorker)
}

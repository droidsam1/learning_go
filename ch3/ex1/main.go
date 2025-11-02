package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	fmt.Println(greetings[:2])  //first two elements
	fmt.Println(greetings[1:4]) //second, third and fourth variables
	fmt.Println(greetings[3:5]) //fourth and fifth variables
}

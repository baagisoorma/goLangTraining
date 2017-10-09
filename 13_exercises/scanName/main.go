package main

import (
	"fmt"
)

func main() {
	var a string

	fmt.Print("Hello, what is your name? ")
	fmt.Scan(&a)

	fmt.Printf("Welcome to Go, %s!\n", a)
}
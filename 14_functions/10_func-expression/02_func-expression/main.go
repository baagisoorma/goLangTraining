package main

import "fmt"

func main() {
	/* The func has no name, it gets assigned to a variable
	 */
	greeting := func(){
		fmt.Println("Sat Sri Akal, World!")
	}

	greeting()
}
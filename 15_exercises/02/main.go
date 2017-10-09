package main

import "fmt"

func main(){
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	divider := func(a int)(int, bool) {
		return a/2, a % 2 == 0
 		}

	fmt.Println(divider(number))
}

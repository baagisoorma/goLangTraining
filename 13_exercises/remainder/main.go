package main

import "fmt"

func main(){
	var big int
	var small int

	fmt.Print("Input a large number: ")
	fmt.Scan(&big)
	fmt.Print("Input a smaller number: ")
	fmt.Scan(&small)

	fmt.Printf("The remainder of %d divided by %d is %d\n", big, small, big % small)
}

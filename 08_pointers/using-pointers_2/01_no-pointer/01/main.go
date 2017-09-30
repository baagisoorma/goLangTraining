package main

import "fmt"

func zero(x int){
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)
}

// x is expected to be 5 still

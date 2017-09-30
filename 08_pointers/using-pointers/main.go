package main

import "fmt"

func main(){
	a := 43
	fmt.Println(a)

	var b *int = &a

	*b = 41
	fmt.Println(a)
}

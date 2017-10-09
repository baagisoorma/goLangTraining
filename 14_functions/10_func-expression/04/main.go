package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Sat Sri Akal, World!"
	}
}

func main(){
	greet := makeGreeter()
	fmt.Println(greet())
}

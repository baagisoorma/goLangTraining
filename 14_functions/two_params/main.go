package main

import "fmt"

func main(){
	greet("Savi", "Kooner")
}

func greet(fname string, lname string){
	fmt.Printf("Hello, Ms. %s. Do you go by \"%s\"?\n", lname, fname)
}

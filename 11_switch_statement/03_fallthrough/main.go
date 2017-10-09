package main

import "fmt"

func main(){
	switch "Savi" {
	case "Raj":
		fmt.Println("Raj!")
	case "Savi":
		fmt.Println("Savi!")
		fallthrough
	case "Prince":
		fmt.Println("Prince!")
		fallthrough
	case "Love":
		fmt.Println("Love!")
	}
}

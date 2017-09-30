package main

import "fmt"

func main(){
	myFriendsName := "Savi"

	switch {
	case len(myFriendsName) == 0:
		fmt.Println("What is my friend's name???")
	case len(myFriendsName) > 6:
		fmt.Println("Yayyy, my friend's name is more than 6 letters")
	case len(myFriendsName) < 6:
		fmt.Println("Yayyy, my friend's name is less than 6 letters")
	default:
		fmt.Println("Yayyy, my friend's name is exactly 6 letters!")
	}
}

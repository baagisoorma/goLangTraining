package main

import "fmt"

type Contact struct {
	greeting	string
	name		string
}

func printContact(x Contact){
	fmt.Println(x.greeting,", ", x.name)
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("other")
	}
}

func main(){
	SwitchOnType(7)
	SwitchOnType("Prince")
	var t = Contact{"Sat Sri Akal", "Prince"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
	printContact(t)
}

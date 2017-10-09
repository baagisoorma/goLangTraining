package main

import "fmt"

func main(){
	fmt.Println(greet("Savi ", "Sukh"))
}

func greet(aname string, bname string) string {
	return fmt.Sprint(aname, bname)
}


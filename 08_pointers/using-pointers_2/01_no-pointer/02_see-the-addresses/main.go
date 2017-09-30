package main

import "fmt"

func zero(x int){
	fmt.Println(&x)
	x = 0
}

func main(){
	x := 5

	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
}

/* 	Same. We expected x to be 5 still, however, we just printed addresses to
 	show why
*/

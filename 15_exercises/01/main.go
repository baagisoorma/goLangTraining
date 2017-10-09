package main

import "fmt"

func half(number float64) (float64,bool){
	var halved float64 = number/2
	return halved, int(number) % 2 == 0

}
func main(){
	var number int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&number)
	halved, even := half(float64(number))

	if even == true {
		fmt.Printf("Half of %d is %f and %d is even\n", number, halved, number)
	} else {
		fmt.Printf("Half of %d is %f and %d is odd\n", number, halved, number)
	}

}

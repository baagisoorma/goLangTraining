package main

import "fmt"

func main(){
	data := []float64{2,2,2,2,2}

	n := average(data...)
	//Even though we have a list of 5 items, the ... tells program
	//to pass in each item one at a time
	fmt.Println(n)
}

//Also notice that data is of type slice but average takes float64, so it wouldn't
//work without the ... above
func average(sf ...float64) float64{
	var total float64

	for _,v := range sf{
		total += v
	}

	return total/float64(len(sf))
}
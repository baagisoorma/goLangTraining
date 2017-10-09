package main

import "fmt"

func biggest(nums...int){
	biggest := 0
	for _,i := range nums{
		if i > biggest{
			biggest = i
		}
	}

	fmt.Println(biggest)
}

func main(){
	list_a := []int{1,2,16,4,0}
	list_b := []int{4,7,1,8,19,0}

	biggest(list_a...)
	biggest(list_b...)
}

package main

import "fmt"

func foo(nums...int){
	sum := 0
	for _,i := range nums{
		sum += i
	}
	fmt.Println(sum)
}

func main(){
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
	foo()
}

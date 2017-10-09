package main

import "fmt"

func main(){
	fmt.Println(combineNames("Savi ", "Sukh"))
}

func combineNames(aname string, bname string) (s string) {
	s = fmt.Sprint(aname, bname)

	return
}

package main

import "fmt"

func main(){
	fmt.Println(combineNames("-Savi-", "-Sukh-"))
}

func combineNames(aname string, bname string) (string, string){
	return fmt.Sprint(aname, bname), fmt.Sprint(bname, aname)
}
1
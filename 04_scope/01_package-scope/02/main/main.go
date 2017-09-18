package main

import (
	"fmt"
	"github.com/golangTraining/04_scope/01_package-scope/02/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}

package main

import (
	"fmt"

	"github.com/roguefactories/golang_beginner/02_The_Init_Function/pkg1"
	"github.com/roguefactories/golang_beginner/02_The_Init_Function/pkg2"
)

func init() {
	fmt.Println("This is init in main.")
}

func main() {
	fmt.Println("This is main.")
	pkg1.Pkg1()
	pkg1.Pkg1_1()
	pkg2.Pkg2()
}

package pkg2

import (
	"fmt"

	"github.com/roguefactories/golang_beginner/02_The_Init_Function/pkg3"
)

func init() {
	fmt.Println("This is init in pkg2.")
}

func Pkg2() {
	fmt.Println("This is pkg2.")
	pkg3.Pkg3()
}

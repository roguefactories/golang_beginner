package pkg1

import (
	"fmt"

	"github.com/roguefactories/golang_beginner/02_init/pkg3"
)

func init() {
	fmt.Println("This is init in pkg1.")
}

func Pkg1() {
	fmt.Println("This is pkg1.")
	pkg3.Pkg3()
}

package main

import "fmt"

type Dog struct {
	Name      string
	Age       int
	Height    float64
	Trainable bool
}

func main() {
	roguefactories := Dog{Name: "RogueFactories", Age: 23, Height: 123.456789, Trainable: true}

	fmt.Printf("%v\n", roguefactories)
	fmt.Printf("%+v\n", roguefactories)
	fmt.Printf("%#v\n", roguefactories)
	fmt.Printf("%T\n", roguefactories)

	fmt.Printf("%t\n", roguefactories.Trainable)

	fmt.Printf("%b\n", roguefactories.Age)
	fmt.Printf("%c\n", roguefactories.Age)
	fmt.Printf("%d\n", roguefactories.Age)
	fmt.Printf("%o\n", roguefactories.Age)
	fmt.Printf("%O\n", roguefactories.Age)
	fmt.Printf("%q\n", roguefactories.Age)
	fmt.Printf("%x\n", roguefactories.Age)
	fmt.Printf("%X\n", roguefactories.Age)
	fmt.Printf("%U\n", roguefactories.Age)

	fmt.Printf("%b\n", roguefactories.Height)
	fmt.Printf("%e\n", roguefactories.Height)
	fmt.Printf("%E\n", roguefactories.Height)
	fmt.Printf("%f\n", roguefactories.Height*10000)
	fmt.Printf("%F\n", roguefactories.Height*1000)
	fmt.Printf("%g\n", roguefactories.Height*10000)
	fmt.Printf("%G\n", roguefactories.Height*1000)

	fmt.Printf("%s\n", roguefactories.Name)
	fmt.Printf("%q\n", roguefactories.Name)
}

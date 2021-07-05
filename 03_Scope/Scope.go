package main

import "fmt"

// global variable declaration
var x int = 10
var y int = 20

func main() {
	// Local variable declaration
	var a, b, c int
	a = 1
	b = 2
	c = a + b

	fmt.Printf("%d + %d = %d\n", a, b, c)
	fmt.Printf("x = %d and y = %d in main()\n", x, y)

	var x int = 1
	s := sum(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, s)
}

func sum(x, y int) int {
	fmt.Printf("x = %d and y = %d in sum()\n", x, y)

	return x + y
}

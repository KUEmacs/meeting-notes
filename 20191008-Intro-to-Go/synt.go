package main

import "fmt"

func main() {
	var a int = 10
	var b = 37
	c := 42
	fmt.Printf("Types for a, b, c: %T, %T, %T\n", a, b, c)

	// Casting
	var fa float64 = 10
	var fb = float64(10)
	fc, fd := 10.0, 12
	fmt.Printf("Types for fa, fb, fc, fd: %T, %T, %T %T\n", fa, fb, fc, fd)
}

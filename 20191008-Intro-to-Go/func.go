package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var a int = 42
	var b int = 13
	fmt.Println(add(a, b))
}

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Equivalent to while true in C
	for {
		fmt.Println("hello, world")
	}
}

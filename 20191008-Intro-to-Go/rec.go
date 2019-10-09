package main

import "fmt"

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	for i := int(0); i < 10; i++ {
		fmt.Println(fib(i))
	}
}

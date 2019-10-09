package main

import "fmt"

func main() {
	defer fmt.Println("exiting the main function...")

	fmt.Println("my print")
}

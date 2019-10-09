package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	// Accessing fields
	v.X = 42
	fmt.Println(v)
	// Pointing to struct
	p := &v // Type Vertex
	p.Y = 100
	fmt.Println(v)
}

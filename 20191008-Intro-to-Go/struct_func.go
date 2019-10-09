package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) Distance() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func main() {
	v1 := Vertex{2, 3} // has type Vertex
	fmt.Println(v1)
	fmt.Println(v1.Distance())
}

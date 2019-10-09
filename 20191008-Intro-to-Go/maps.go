package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	m["me"] = Vertex{42, 13}
	fmt.Printf("%v is a %T\n", m["Google"], m["Google"])
	fmt.Println(m)
}

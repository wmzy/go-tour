package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = map[string]Vertex{
		"Bell Labs": {40.342, -74.23434},
		"Google":    {30.342, -174.23434},
	}
	fmt.Println(m)
}

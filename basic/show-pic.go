package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, 0, dy)
	for i := 0; i < dy; i++ {
		row := make([]uint8, dx)
		row[i%dx] = uint8(i)
		p = append(p, row)
	}

	return p
}

func main() {
	pic.Show(Pic)
}

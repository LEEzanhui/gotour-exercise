package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		var one []uint8
		for j := 0; j < dx; j++ {
			one = append(one, (uint8)(i*j))
		}
		pic = append(pic, one)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}

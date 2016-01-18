package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			ret[i][j] = uint8(i ^ j)
		}
	}
	return ret
}

// 参考 https://tour.go-zh.org/moretypes/15
func main() {
	// nice choice (x+y)/2、x*y 和 x^y 
	pic.Show(Pic)
}

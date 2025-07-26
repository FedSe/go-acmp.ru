package main
import . "fmt"
func main() {
	var (
		Y, X, g, v, w, y, i int
		a = []int{1, 0, -1, 0, 0, 1, 0, -1}
		o = 1
		p [102][102]int
	)

	Scan(&y, &i, &Y, &X)

	for 0 < y {
		for
		x := 1
		x <= i
		{
			p[y][x] = 1
		x++
		}
	y--
	}

	p[Y][X] = 2
	X = 1
	Y = 1
	for p[Y][X] != 2 {
		if p[Y][X] < 1 {
			X -= o
			Y -= w
			g = (g + 5) % 4
			o = a[g]
			w = a[g+4]
			X += o
			Y += w
		}
		if p[Y][X] != 2 {
			v++
			p[Y][X] = 0
			X += o
			Y += w
		}
	}
	v++

	Print(v)
}
package main
import . "fmt"
type T [2]int

var (
	q, r, k string
	A       [8][8]int
	O       = map[any]int{}
	C, l    int
	K       = []int{-2, -1, -2, 1, -1, -2, -1, 2, 1, -2, 1, 2, 2, -1, 2, 1}
)

func F(s string) (int, int) {
	c := int(s[0] - 65)
	R := 8 - int(s[1]-48)
	O[T{R, c}] = 1
	return R, c
}
func main() {
	Scan(&q, &r, &k)
	a, b := F(q)
	c, e := F(r)
	f, g := F(k)
	for i, d := range []T{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
		x := a + d[0]
		y := b + d[1]
		for x >= 0 && x < 8 && y >= 0 && y < 8 {
			A[x][y] = 1
			x += d[0]
			y += d[1]
		}
		x = f + K[i*2]
		y = g + K[i*2+1]
		if x >= 0 && x < 8 && y >= 0 && y < 8 {
			A[x][y] = 1
		}
		if i < 4 {
			x = c + d[0]
			y = e + d[1]
			for x >= 0 && x < 8 && y >= 0 && y < 8 {
				A[x][y] = 1
				x += d[0]
				y += d[1]
			}
		}
	}

	for l < 8 {
		j := 0
		for j < 8 {
			if O[T{l, j}] < A[l][j] {
				C++
			}
			j++
		}
		l++
	}
	Print(C)
}
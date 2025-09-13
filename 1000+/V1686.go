package main
import . "fmt"
func main() {
	var A, B, X, Y, q int

	Scan(&A, &B, &X, &Y)
	for _, r := range []int{A + X, A - X, B + Y, B - Y} {
		x := r - A
		if x < 0 {
			x = -x
		}
		y := r - B
		if y < 0 {
			y = -y
		}
		if x == X && y == Y {
			q = r
		}
	}

	Print(q)
}
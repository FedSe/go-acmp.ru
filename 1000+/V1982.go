package main
import . "fmt"
func main() {
	var M, R, A int
	N := 100

	Scan(&M)
	for N < 1e3 {
		b := N / 10 % 10
		x := N / 100 * b
		y := b * (N % 10)
		if x > y {
			x, y = y, x
		}
		if x+y < 1 {
			R = 0
		} else if x < 1 {
			R = y
		} else {
			R = x
			t := y
			for t > 0 {
				R *= 10
				t /= 10
			}
			R += y
		}
		if R == M {
			A = N
			break
		}
		N++
	}

	Print(A)
}
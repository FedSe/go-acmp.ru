package main
import . "fmt"
func M(a, b int) int {
	if a > b {
		b = a
	}
	return b
}
func main() {
	var z, x, c, v, A, B, C, D, X, Y, p int
	m := 1

	Scan(&z, &x, &c, &v, &A, &B, &C, &D, &X, &Y)
	w := X*X + Y*Y
	if w <= x*x {
		if w >= z*z && w <= x*x {
			m = 2
		}
		if w >= c*c && w <= v*v {
			m = 3
		}
		if Y == 0 {
			if X > 0 {
				B = A
				C = D
			}
			p = M(B, C)
		}
		if X == 0 {
			if Y > 0 {
				C = A
				D = B
			}
			p = M(C, D)
		}
		if X == 0 && Y == 0 {
			p = M(M(A, B), M(C, D))
		}
		if X > 0 && Y < 0 {
			p = D
		}
		if X < 0 && Y < 0 {
			p = C
		}
		if X < 0 && Y > 0 {
			p = B
		}
		if X > 0 && Y > 0 {
			p = A
		}
	}

	Print(p * m)
}
package main
import . "fmt"
func main() {
	var A, B, C [3]float64
	Scan(&A[0], &A[1], &A[2], &B[0], &B[1], &B[2], &C[0], &C[1], &C[2])

	F := func(u, v [3]float64) (float64, float64) {
		d := u[0]*v[1] - v[0]*u[1]
		return (u[2]*v[1] - v[2]*u[1]) / d, (u[0]*v[2] - v[0]*u[2]) / d
	}

	a, b := F(A, B)
	c, d := F(B, C)
	e, f := F(C, A)

	a = (c-a)*(f-b) - (e-a)*(d-b)
	if a < 0 {
		a = -a
	}

	Print(a / 2)
}
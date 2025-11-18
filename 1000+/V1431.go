package main
import . "fmt"
func main() {
	var N, A, B, C, D int
	Scan(&N, &A, &B, &C, &D)

	A += B + C + D
	A -= 3 * N
	if A < 0 {
		A = 0
	}

	Print(A)
}
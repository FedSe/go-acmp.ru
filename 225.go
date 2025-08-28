package main
import . "fmt"
func main() {
	var (
		n, i, l    int
		A, B, C, D [6e3]int
	)

	Scan(&n)
	for l < n {
		Scan(&A[l], &B[l], &C[l])
		l++
	}

	for i < n {
		i++
		D[i] = D[i-1] + A[i-1]
		if i > 1 {
			l = D[i-2] + B[i-2]
			if l < D[i] {
				D[i] = l
			}
		}
		if i > 2 {
			l = D[i-3] + C[i-3]
			if l < D[i] {
				D[i] = l
			}
		}
	}

	Print(D[n])
}
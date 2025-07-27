package main
import . "fmt"
func main() {
	var (
		x                   [1001]int
		j, i, N, A, B, C, D int
	)

	Scan(&N, &A, &B, &C, &D)

	for i < N {
		i++
		x[i] = i
	}

	B += A
	for A <= B/2 {
		i = x[A]
		x[A] = x[B-A]
		x[B-A] = i
		A++
	}

	D += C
	for C <= D/2 {
		i = x[C]
		x[C] = x[D-C]
		x[D-C] = i
		C++
	}

	for j < N {
		j++
		Println(x[j])
	}
}
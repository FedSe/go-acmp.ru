package main
import (
	. "fmt"
	. "math"
	. "sort"
)
func main() {
	a := "Yes"
	s := a
	n := 0
	w := 0

	Scan(&s)
	A := []byte(s)
	for w < 2 {
		Slice(A, func(i, j int) bool {
			if w < 1 {
				i, j = j, i
			}
			return A[i] > A[j]
		})

		Sscan(string(A), &n)
		i := 1
		for i < int(Sqrt(float64(n))) {
			i++
			if n%i < 1 {
				a = "No"
			}
		}
		w++
	}

	Print(a)
}
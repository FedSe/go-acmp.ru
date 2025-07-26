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
	Scan(&s)
	A:=[]byte(s)

	for
	w := 0
	w < 2
	{
		Slice(A, func(i, j int) bool {
			if w < 1 {
				return A[i] < A[j]
			}
			return A[i] > A[j]
		})

		Sscan(string(A), &n)

		for
		i := 1
		i < int(Sqrt(float64(n)))
		{
		i++
			if n % i < 1 {
				a = "No"
			}
		}

	w++
	}

	Print(a)
}
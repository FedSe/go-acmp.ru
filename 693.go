package main
import (
	. "fmt"
	. "sort"
	. "strings"
)
func main() {
	s := "No"
	a := s
	b := s
	i := 0

	for i < 2 {
		Scan(&a)
		A := []byte(ToLower(a))
		Slice(A, func(i, j int) bool {
			return A[i] < A[j]
		})
		a = string(A)

		if a == b {
			s = "Yes"
		}

		b = a
		i++
	}

	Print(s)
}
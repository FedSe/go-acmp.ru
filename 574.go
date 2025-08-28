package main
import (
	. "fmt"
	. "sort"
)

func main() {
	s := "NO"
	a := s
	b := s
	i := 0

	for i < 2 {
		Scan(&a)
		A:=[]byte(a)
		Slice(A, func(i, j int) bool {
			return A[i] < A[j]
		})
		a = string(A)

		if a == b {
			s = "YES"
		}

		b = a
	i++
	}

	Print(s)
}
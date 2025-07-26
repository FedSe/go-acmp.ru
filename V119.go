package main
import (. "fmt"
. "sort"
)
func main() {
	var (i, n, a, b, c, j int
	A []int
)
	Scan(&n)

	for i < n {
		Scan(&a, &b, &c)
		A = append(A, a*3600 + b*60 + c)
	i++
	}

	Ints(A)

	for j < n {
		Println(A[j]/3600, A[j]/60%60, A[j]%60)
	j++
	}
}
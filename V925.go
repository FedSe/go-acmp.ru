package main
import (
   . "fmt"
   . "sort"
)
func main() {
	var k, n, a, b, c int
	Scan(&k, &n, &a, &b, &c)

	A := []int{a, b, c}
	Ints(A)

	n = a+b+c-2*n
	if n < 0 { n = 0 }
	if k > 1 { n = A[0] }

	Print(n)
}
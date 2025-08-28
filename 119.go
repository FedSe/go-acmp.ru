package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		n, a, b, c int
		A          []int
	)
    
	Scan(&n)
	for 0 < n {
		Scan(&a, &b, &c)
		A = append(A, a*3600+b*60+c)
		n--
	}

	Ints(A)

	for _, v := range A {
		Println(v/3600, v/60%60, v%60)
	}
}
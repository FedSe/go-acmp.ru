package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		v []int
		n, a int
		P = Println
	)

	Scan(&n)
	for 0 < n {
		Scan(&a)
		if a&1 < a/64&1 {
			v = append(v, a)
		}
		n--
	}
	Ints(v)

	P(len(v))
	for _, i := range v {
		P(i)
	}
}
package main
import . "fmt"
var a [1001]int

func F(n int) int {
	if n > 2 {
		if a[n] > 0 {
			return a[n]
		}
		k := F(n/2) + F(n-2)
		a[n] = k
		n = k
	}
	return n
}

func main() {
	n := 0
	Scan(&n)
	Print(F(n))
}
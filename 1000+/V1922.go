package main
import . "fmt"
func main() {
	n := 0
	k := 0
	a := 6

	Scan(&n, &k)
	if k < 2*n {
		a = 5
	}
	if k == n {
		a = 3
	}
	if k < n {
		a = 4
	}

	Print(a)
}
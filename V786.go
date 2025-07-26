package main
import . "fmt"
func main() {
	n := 0
	a := 0
	t := 2

	Scan(&n)
	for t < n {
		t *= 2
	}
	if n > 2 && t != n {
		a = n - t/2
	}

	Print(a)
}
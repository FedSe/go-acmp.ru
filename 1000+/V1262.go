package main
import . "fmt"
func main() {
	var x, n, t int

	Scan(&x, &n, &t)
	t = 240 / t
	if n-x < t {
		t = n - x
	}

	Print(x + t)
}
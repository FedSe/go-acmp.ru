package main
import . "fmt"
func main() {
	var n, t, w, i, a int

	Scan(&n)
	for i < n {
		Scan(&a)
		t += a
		w += a * (n - 1 - i)
		i++
	}

	Scan(&a)
	Print((a - t) / w)
}
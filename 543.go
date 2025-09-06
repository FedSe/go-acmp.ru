package main
import . "fmt"
func main() {
	var n, w, d, p int

	Scan(&n, &w, &d, &p)
	w = w*(n-1)*n/2 - p
	if w > 0 {
		n = w / d
	}

	Print(n)
}
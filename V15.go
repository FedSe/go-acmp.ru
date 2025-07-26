package main
import . "fmt"
func main() {
	var n, a, w int
	Scan(&n)
	n*=n
	for 0 < n {
		Scan(&a)
		w += a
	n--
	}

	Print(w / 2)
}
package main
import . "fmt"
func main() {
	a := 0
	n := 0

	Scan(&a, &n)
	if a > 9*n {
		n = 1
		a = -1
	}
	for 0 < n {
		d := 9
		if d > a {
			d = a
		}
		Print(d)
		a -= d
		n--
	}
}
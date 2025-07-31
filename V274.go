package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	for n > 0 {
		var (
			v = `NO
`
			a    = v
			b    = v
			x, y [58]int
		)

		Scan(&a, &b)
		for _, c := range a {
			x[c] = 1
		}
		for _, c := range b {
			y[c] = 1
		}
		if x == y {
			v = `YES
`
		}
		Print(v)
		n--
	}
}
package main
import . "fmt"
func main() {
	var a, b, c int
	m := -1

	Scan(&a, &b, &c)
	for _, x := range [][]int{{a, b, c}, {a, c, b}, {b, a, c}, {b, c, a}, {c, a, b}, {c, b, a}} {
		if x[0] > 0 && x[2]&1 < 1 {
			q := x[0]*100 + x[1]*10 + x[2]
			if q > m {
				m = q
			}
		}
	}

	if m < 0 {
		Print("Impossible")
		return
	}

	Print(m)
}
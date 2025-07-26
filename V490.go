package main
import . "fmt"
func main() {
	var p, m, y, i, d, v, t int
	Scanf("%d.%d.%d\n%d.%d.%d", &p, &m, &y, &d, &v, &t)
	h := [13]int{ -28, 3, 0, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3 }

	for p != d || m != v || y != t {
		p++
		if p > h[m] + 28 {
			p = 1
			m++
		}
		if m > 12 {
			m = 1
			y++
		}
	i++
	}

	Print(i)

}
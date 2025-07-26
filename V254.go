package main
import . "fmt"
func main() {
	var (
		n, m, i, f, t int
		p [201]int
		o [5000]int
	)

	Scan(&n)

	for t < n {
		Scan(&o[t])
	t++
	}

	Scan(&m)

	for f < 201 {
		p[f] = f
	f++
	}

	for 0 < m {
		Scan(&f, &t)
		p[f] = t
	m--
	}

	for i < n {
		Print(p[o[i]], " ")
	i++
	}

}
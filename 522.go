package main
import . "fmt"
type M map[int]int
func main() {
	var (
		n, m, i, t int
		a          = M{}
		b          = M{}
	)

	Scan(&n, &m)
	for i < n+m {
		Scan(&t)
		if i < n {
			a[t] = 1
		} else {
			b[t] = 1
		}
		i++
	}

	i = 1
	for k := range a {
		if b[k] < 1 || len(a) != len(b) {
			i = 0
		}
	}

	Print(i)
}
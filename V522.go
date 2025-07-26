package main
import . "fmt"
func main() {
	var (
		n, m, i, t int
		a = map[int]bool{}
		b = map[int]bool{}
	)
	Scan(&n, &m)

	for i < n+m {
		Scan(&t)
		if i < n {
			a[t] = 1>0
		} else {
			b[t] = 1>0
		}
	i++
	}

	i = 1
	for k := range a {
		if !b[k] || len(a) != len(b) {
			i = 0
			break
		}
	}

	Print(i)
}
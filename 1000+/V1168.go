package main
import . "fmt"
type S struct {
	f string
	b int
}

func main() {
	var (
		a [1001]S
		n, m, x, y, z, i int
	)

	Scan(&n, &m)
	for i < n {
		Scan(&a[i].f, &x, &y, &z)
		a[i].b = x + y + z
		if i > 0 && (a[i].b > a[i-1].b || (a[i].b == a[i-1].b && a[i].f > a[i-1].f)) {
			c := a[i]
			j := i-1
			for j >= 0 && (c.b > a[j].b || (c.b == a[j].b && c.f > a[j].f)) {
				a[j+1] = a[j]
				j--
			}
			a[j+1] = c
		}
	i++
	}

	m--
	Println(a[m].f, a[m].b)
}
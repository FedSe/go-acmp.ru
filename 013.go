package main
import . "fmt"
func main() {
	n := 0
	m := 0
	a := ""
	b := a

	Scan(&a, &b)

	for i, v := range a {
		for j, w := range b {
			if v == w {
				if i == j {
					n++
					m--
				}
				m++
			}
		}
	}

	Print(n, m)
}
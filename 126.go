package main
import . "fmt"
func main() {
	var (
		n, i, s int
		a       [100][100]int
	)

	Scan(&n)
	for i < n {
		j := 0
		for j < n {
			Scan(&a[i][j])
			j++
		}
		i++
	}

	i = 3e3
	for s < n {
		v := s + 1
		for v < n {
			w := v + 1
			for w < n {
				t := a[s][v] + a[v][w] + a[w][s]
				if t < i {
					i = t
				}
				w++
			}
			v++
		}
		s++
	}

	Print(i)
}
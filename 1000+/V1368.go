package main
import . "fmt"
func main() {
	var (
		a             [51][51]int
		n, s, t, i, k int
		F             = 1 << 30
	)

	Scan(&n, &s, &t)
	s--
	t--

	for i < n {
		j := 0
		for j < n {
			Scan(&a[i][j])
			if a[i][j] < 0 {
				a[i][j] = F
			}
			j++
		}
		i++
	}

	for k < n {
		i = 0
		for i < n {
			j := 0
			for j < n {
				f := a[i][k] + a[k][j]
				if a[i][j] > f {
					a[i][j] = f
				}
				j++
			}
			i++
		}
		k++
	}

	i = a[s][t]
	if i >= F {
		i = -1
	}
	Print(i)
}
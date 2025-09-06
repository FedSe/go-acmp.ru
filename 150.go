package main
import . "fmt"
func main() {
	var (
		n, s, a, l, i int
		F             [101][101]int
	)

	Scan(&n, &s)
	for i < n {
		i++
		j := 0
		for j < n {
			j++
			Scan(&F[i][j])
		}
	}

	for l < n {
		l++
		i = 0
		for i < n {
			j := 0
			i++
			for j < n {
				j++
				if F[i][l]*F[l][j] > 0 {
					F[i][j] = 1
				}
			}
		}
	}

	for n > 0 {
		if n != s && F[s][n] > 0 {
			a++
		}
		n--
	}

	Print(a)
}
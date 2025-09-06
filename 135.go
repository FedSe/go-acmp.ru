package main
import . "fmt"
func main() {
	var (
		a             [101][101]int
		n, i, k, l, j int
	)

	Scan(&n)
	for i < n {
		j = 0
		for j < n {
			Scan(&a[i][j])
			j++
		}
		i++
	}

	for k < n {
		i = 0
		for i < n {
			j = 0
			for j < n {
				v := a[i][k] + a[k][j]
				if a[i][j] > v {
					a[i][j] = v
				}
				j++
			}
			i++
		}
		k++
	}

	for l < n {
		j = 0
		for j < n {
			Println(a[l][j])
			j++
		}
		l++
	}
}
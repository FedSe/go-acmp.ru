package main
import . "fmt"
func main() {
	var (
		d          [251]int
		N, c, i, j int
	)

	Scan(&N)
	for j < N {
		j++
		d[j] = 1e7
	}

	for i < N {
		j = i
		for j < N {
			j++
			Scan(&c)
			if d[i]+c < d[j] {
				d[j] = d[i] + c
			}
		}
		i++
	}

	Print(d[N])
}
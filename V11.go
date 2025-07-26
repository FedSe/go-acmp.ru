package main
import . "fmt"
func main() {
	var (
		m [300][99]int
		k, n, i int
	)

	Scan(&k, &n)
	for i < n {
		if i < k {
			m[i][0] = 1
		}
		j := i-k
		if j < 0 { j = 0 }
		for j < i {
			t := 0
			w := 0
			for t < 99 {
				w = m[i][t] + m[j][t] + w / 10
				m[i][t] = w % 10
			t++
			}
		j++
		}
	i++
	}

	n--
	i = 98
	for m[n][i] < 1 {
		i--
	}

	for i >= 0 {
		Print(m[n][i])
	i--
	}
}
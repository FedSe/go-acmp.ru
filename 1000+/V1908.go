package main
import . "fmt"
func main() {
	var N, M, K, i int
	s := `Possible
`
	Scan(&N, &M, &K)
	j := N
	if j > M {
		j = M
	}
	if j < K {
		s = "Impossible"
		N = 0
	}
	for i < N {
		i++
		j = 0
		for j < M {
			j++
			w := "."
			if i == j && i <= K {
				w = "#"
			}
			s += w
		}
		s += `
`
	}

	Print(s)
}
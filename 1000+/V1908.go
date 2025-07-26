package main
import . "fmt"
func main() {
	var N, M, K, i int
	s := "Possible\n"
	Scan(&N, &M, &K)
	a := N
	if a > M { a = M }
	if a < K {
		s = "Impossible"
	} else {
		for i < N {
		i++
			j := 0
			for j < M {
			j++
				w := "."
				if i == j && i <= K {
					w = "#"
				}
				s += w
				
			}
			s += "\n"
		}
	}
	Print(s)
}
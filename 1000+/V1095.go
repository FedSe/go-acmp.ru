package main
import . "fmt"
func main() {
	var N, M, K, i int

	Scan(&N, &M, &K)
	for i < N*K {
		j := 0
		for j < M*K {
			a := "."
			if (N-1-i/K+j/K)&1 < 1 {
				a = "X"
			}
			Print(a)
			j++
		}
		Print(`
`)
		i++
	}
}
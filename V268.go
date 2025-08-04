package main
import . "fmt"
func main() {
	var N, K, i int
	S := ""
	Scan(&N, &K, &S)

	a := make([]int, N)
	b := make([]int, N-1)
	c := N
	for i < N-1 {
		if S[i] != S[i+1] {
			b[i] = 1
		}
		if b[i] <= K {
			c++
		}
		i++
	}

	i = 1
	for i < N {
		i++
		u := make([]int, N-i)
		l := 0
		for l < len(u) {
			u[l] = a[l+1]
			if S[l] != S[l+i] {
				u[l]++
			}
			if u[l] <= K {
				c++
			}
			l++
		}
		a = b
		b = u
	}

	Print(c)
}
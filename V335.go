package main
import . "fmt"
const M = 1e9 + 9
type T [100]int
func F(n int) int {
	if n < 3 {
		return n - 1
	}
	if n%2 < 1 {
		return 0
	}
	i := 3
	for i*i <= n {
		if n%i < 1 {
			return 0
		}
		i += 2
	}
	return 1
}

func main() {
	var (
		p    []int
		N, t int
		i    = 4
		n    = 100
		m    T
	)

	Scan(&N)
	for n < 1e3 {
		if F(n) > 0 {
			p = append(p, n)
		}
		n++
	}

	if N == 3 {
		Print(len(p))
		return
	}

	for _, v := range p {
		b := (v / 10) % 10
		e := b*10 + v%10
		m[e] = (m[e] + 1) % M
	}

	for i <= N {
		var w T
		q := 0
		for q < 100 {
			a := q / 10
			b := q % 10
			if m[q] > 0 {
				c := 0
				for c < 10 {
					if a > 0 && F(a*100+b*10+c) > 0 {
						e := b*10 + c
						w[e] = (w[e] + m[q]) % M
					}
					c++
				}
			}
			q++
		}
		m = w
		i++
	}

	for _, v := range m {
		t = (t + v) % M
	}
	Print(t)
}
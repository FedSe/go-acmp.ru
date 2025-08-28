package main
import . "fmt"
type A map[int]int
func main() {
	var (
		m, n, t, j int
		i          = 1
		q          = A{}
	)

	Scan(&m, &n)
	if m > n {
		m, n = n, m
	}
	e := 1 << m
	for j < e {
		q[j] = 1
		j++
	}
	for i < n {
		o := A{}
		p := 0
		for p < e {
			h := 0
			for h < e {
				j = 0
				for j < m-1 {
					a := p >> j & 1
					c := h >> j & 1
					j++
					b := p >> j & 1
					if a == b && b == c && c == h>>j&1 {
						goto A
					}
				}
				o[h] += q[p]
			A:
				h++
			}
			p++
		}
		q = o
		i++
	}

	for _, v := range q {
		t += v
	}

	Print(t)
}
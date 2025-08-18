package main
import . "fmt"
func main() {
	var (
		a, b, c, d, k, i int
		s                = map[int]int{}
		h                []int
	)

	Scan(&a, &b, &c, &d, &k)
	n := a
	for i < k {
		p, o := s[n]
		if o {
			i = p + (k-p)%(i-p)
			if k < p {
				i = k
			}
			n = h[i]
			break
		}
		s[n] = i
		h = append(h, n)
		if n*b < c {
			n = 0
			break
		}
		n = n*b - c
		if n > d {
			n = d
		}
		i++
	}

	Print(n)
}
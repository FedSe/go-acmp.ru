package main
import . "fmt"
func main() {
	var (
		o                = [99]int{1}
		k, w, q, i, c, j int
		p, m, A, B, C, D uint32
	)

	for j < 2 {
		p = m
		m = 0
		Scanf(`%d.%d.%d.%d
`, &A, &B, &C, &D)
		m = m<<8 + A
		m = m<<8 + B
		m = m<<8 + C
		m = m<<8 + D
		j++
	}
	Scan(&k)

	p &= m
	m = ^m
	for m > 0 {
		q += int(m & 1)
		m >>= 1
	}
	for p > 0 {
		w += int(p & 1)
		p >>= 1
	}

	for i < q {
		i++
		j = i
		for j > 0 {
			o[j] += o[j-1]
			j--
		}
	}
	i = 1 << q
	for c <= q {
		if (w+c)%k < 1 {
			i -= o[c]
		}
		c++
	}

	Print(i)
}
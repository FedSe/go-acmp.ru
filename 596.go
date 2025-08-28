package main
import . "fmt"
func main() {
	var (
		x, y, r, v, k    [1e4]int
		m                [100]string
		n, f, i, j, l, o int
		u                = -1
		S                = Scan
		P                = Println
	)

	S(&n)
	for i < n {
		a := ""
		S(&a)
		o = -1
		f = 0
		for f <= u {
			if a == m[f] {
				o = f
				break
			}
			f++
		}
		if o == -1 {
			u++
			m[u] = a
			o = u
		}
		v[i] = o
		S(&x[i], &y[i], &r[i])
		i++
	}

	S(&f, &o)
	for l < n {
		i = x[l] - f
		c := y[l] - o
		if i*i+c*c <= r[l]*r[l] {
			k[v[l]]++
		}
		l++
	}

	P(u + 1)
	for j <= u {
		P(m[j], k[j])
		j++
	}
}
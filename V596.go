package main
import . "fmt"
func main() {
	var (
		x, y, r, v, k [10000]int
		m [100]string
		n, f, i, j, l, o int
		u = -1
	)

	Scan(&n)
	for i < n {
		a := ""
		Scan(&a)
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
		Scan(&x[i], &y[i], &r[i])
	i++
	}

	Scan(&f, &o)
	for l < n {
		i = x[l]-f
		c := y[l]-o
		if i*i+c*c <= r[l]*r[l] {
			k[v[l]]++
		}
	l++
	}

	Println(u + 1)
	for j <= u {
		Println(m[j], k[j])
	j++
	}
}
package main
import . "fmt"
func main() {
	var (
		v          [1e4]int
		s          = ""
		n, i, f, l int
		P          = Println
		S          = Scan
	)

	S(&s, &n)
	m := len(s)
	w := m

	for i < n {
		i++
		t := ""
		S(&t)

		h := 0
		j := 0
		for j < w {
			if s[j] != t[j] {
				h++
			}
			j++
		}

		if h <= m {
			if h < m {
				m = h
				f = 0
			}
			v[f] = i
			f++
		}
	}

	P(f)
	for l < f {
		P(v[l])
		l++
	}
}
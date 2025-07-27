package main
import . "fmt"
func main() {
	var (
		v          [9999]int
		s          = ""
		n, i, f, l int
	)
	Scan(&s, &n)
	m := len(s)
	w := m

	for i < n {
		i++
		t := ""
		Scan(&t)

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

	Println(f)
	for l < f {
		Println(v[l])
		l++
	}
}
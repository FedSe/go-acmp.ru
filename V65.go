package main
import . "fmt"
func main() {
	var (
		v [9999]int
		s = ""
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
		for
		j := 0
		j < w
		{
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
		Print(v[l], " ")
	l++
	}

}
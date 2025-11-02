package main
import . "fmt"
var (
	s    = ""
	n, k int
	m    [256]bool
)

func F(l, r int) bool {
	if l+1 == r {
		return m[s[l]]
	}
	e := 1 > 0
	if l+1 < r {
		var (
			v []int
			c = 0
			i = l + 4
			y = s[l+1]
			f = s[l]
		)
		if f == 65 && y == 78 {
			for i < r {
				t := s[i]
				if t == 44>>c {
					v = append(v, i)
				}
				if t < 41 {
					c++
				}
				if t == 41 {
					c--
				}
				i++
			}
			e = e && F(l+4, v[0])
			i = 1
			for i < len(v) {
				e = e && F(v[i-1]+1, v[i])
				i++
			}
			e = e && F(v[len(v)-1]+1, r-1)
		}
		if f == 79 && y == 82 {
			e = !e
			i--
			for i < r {
				t := s[i]
				if t == 44>>c {
					v = append(v, i)
				}
				if t < 41 {
					c++
				}
				if t == 41 {
					c--
				}
				i++
			}
			e = e || F(l+3, v[0])
			i = 1
			for i < len(v) {
				e = e || F(v[i-1]+1, v[i])
				i++
			}
			e = e || F(v[len(v)-1]+1, r-1)
		}
		if f == 78 && y == 79 {
			e = !F(l+4, r-1)
		}
	}
	return e
}

func main() {
	Scan(&s, &n, &k)
	for 0 < n {
		u := `FALSE
`
		for j := range m {
			m[j] = 1 < 0
		}
		j := 0
		for j < k {
			w := s
			o := ' '
			Scanf(`
%c=%s`, &o, &w)
			m[o] = w[0] == 84
			j++
		}
		if F(0, len(s)) {
			u = `TRUE
`
		}
		Print(u)
		n--
	}
}
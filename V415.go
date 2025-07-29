package main
import (
	. "fmt"
	. "strings"
	u "unicode"
)
func main() {
	var (
		a = ""
		b = a
		l = 2
		L = ToLower
		U = u.ToUpper
	)

	Scan(&a, &b)

	m := a + b
	for l > 0 {
		q := L(a)
		i := 0
		for i <= len(a) {
			w := q[:i] + L(b)
			k := 0
			c := 0
			for c < 3 {
				j := Index(w[k:], q)
				if j < 0 {
					break
				}
				k += j

				e := []rune(w)
				e[i] = U(e[i])
				e[k] = U(e[k])
				s := string(e)

				r := len(s)
				t := len(m)
				if r < t || (r == t && s < m) {
					m = s
				}

				k++
				c++
			}
			i++
		}
		l--
		a, b = b, a
	}

	Print(m)
}
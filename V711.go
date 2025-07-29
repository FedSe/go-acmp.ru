package main
import . "fmt"
func main() {
	var (
		n, m, a int
		s       = ""
		c       = s
		p       = 999999
		S       = Scan
	)

	S(&n, &m)
	for 0 < n {
		l := 0
		S(&s)
		j := 0
		for j < m {
			S(&a)
			l += a
			j++
		}
		if l < p {
			p = l
			c = s
		}
		n--
	}

	Print(c)
}
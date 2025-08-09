package main
import . "fmt"
func main() {
	var (
		n, m, a int
		s       = ""
		c       = s
		p       = 1 << 20
		S       = Scan
	)

	S(&n, &m)
	for 0 < n {
		l := 0
		j := 0
		S(&s)
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
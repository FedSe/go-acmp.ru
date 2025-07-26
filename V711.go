package main
import . "fmt"
func main() {
	var (
		n, m, a int
		s = ""
		c = s
		p = 999999
	)
    
	Scan(&n, &m)
	for 0 < n {
		l := 0
		Scan(&s)
		for
		j := 0
		j < m
		{
			Scan(&a)
			l+=a
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
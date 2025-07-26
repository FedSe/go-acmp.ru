package main
import (
	. "fmt"
	. "os" 
	b "bufio"
	. "strconv"
)
func main() {
	var (
		a, c [300000]int
		n, m, i int
		s = b.NewScanner(Stdin)
	)

	Scan(&n, &m)
	s.Split(b.ScanWords)

	for i < n+m {
		s.Scan()
		f, _ := Atoi(s.Text())
		if i < n {
			c[f] = 1
		} else {
			a[i-n] = f
		}
	i++
	}

	for 0 < m {
	m--
		if c[a[m]] > 0 {
			c[a[m]] = 2
		}
	}

	for j, v := range c {
		if v > 1 {
			Print(j, " ")
		}
	}

}
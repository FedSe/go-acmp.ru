package main
import . "fmt"
func main() {
	var (
		n, i int
		v    = 1
		c    [100]string
		s    = " Rumpelstiltskin"
	)

	Scan(&n)
	for i < n {
		Scan(&c[i])
		i++
	}

	i = 1
	for i < n {
		q := c[i-1]
		a := q[len(q)-1]
		if a > 90 {
			a -= 32
		}
		if c[i][0] != a {
			break
		}
		v++
		i++
	}

	if v == n && n&1 > 0 || v != n && v&1 > 0 {
		s = " Shrek"
	}

	Print(v, s)
}
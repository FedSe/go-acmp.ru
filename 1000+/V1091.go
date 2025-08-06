package main
import . "fmt"
func main() {
	var (
		n, i, l, v int
		c          [100]string
		s          = " Rumpelstiltskin"
	)

	Scan(&n)
	for l < n {
		Scan(&c[l])
		l++
	}
	for i < n {
		if i > 0 {
			a := c[i-1][len(c[i-1])-1]
			if a > 90 {
				a -= 32
			}
			if c[i][0] != a {
				break
			}
		}
		v++
		i++
	}

	if (v == n && n%2 > 0) || (v != n && (v+1)&1 < 1) {
		s = " Shrek"
	}
	Print(v, s)
}
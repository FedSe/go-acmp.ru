package main
import . "fmt"
func main() {
	var k, m, r, i int
	n := ""

	Scan(&k, &m, &n)
	for i < len(n) {
		c := n[i]
		d := int(c - 55)
		if c > 47 && c < 58 {
			d = int(c - 48)
		}
		r = (r*k + d) % m
		i++
	}

	Print(r)
}
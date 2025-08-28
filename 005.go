package main
import . "fmt"
func main() {
	var (
		n, c, r int
		s       = "YES"
		a       = ""
		S       = Scan
		P       = Println
	)

	S(&n)
	for n > 0 {
		S(&c)
		if c%2 > 0 {
			P(c)
			r++
		} else {
			a += Sprintln(c)
			r--
		}
		n--
	}

	if 0 < r {
		s = "NO"
	}

	P(a + s)
}
package main
import . "fmt"
func main() {
	var (
		n, s, m, v int
		S          = Scan
		P          = Print
	)

	S(&n)
	if n == 1 {
		P("No")
	} else {
		S(&v)
		i := 1
		for i < n {
			i++
			S(&m)
			s += i * m
		}
		P(`
Yes
`, -s)
		i = 1
		for i < n {
			i++
			P(" ", i*v)
		}
	}
}
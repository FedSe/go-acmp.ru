package main
import . "fmt"
func main() {
	n := 0
	p := 1.
	a := p

	Scan(&n)
	for 0 < n {
		Scan(&a)
		p = p*a + (1-p)*(1-a)
	n--
	}

	Print(p)
}
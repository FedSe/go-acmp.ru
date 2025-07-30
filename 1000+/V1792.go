package main
import . "fmt"
func main() {
	var a, b, c, d int

	Scan(&a, &b, &c, &d)

	m := a*b + c*d
	n := a*c + b*d
	if n > m {
		m = n
	}
	n = a*d + b*c
	if n > m {
		m = n
	}

	Print(m)
}
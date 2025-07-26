package main
import . "fmt"
func main() {
	x := 1
	m := 0
	a := 0
	Scan(&m)
	for x > 0 {
		Scan(&x)
		if x > m {
			a = m
			m = x
		} else if x > a {
			a = x
		}
	}
	Print(a)
}
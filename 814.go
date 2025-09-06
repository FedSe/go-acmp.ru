package main
import . "fmt"
func main() {
	x := 0
	c := 0
	a := 2

	Scan(&x)
	for a <= x/2 {
		p := 1
		q := 1
		i := 2
		for i*i <= a {
			if a%i < 1 {
				p = 0
			}
			i++
		}
		i = 2
		for i*i <= x-a {
			if (x-a)%i < 1 {
				q = 0
			}
			i++
		}
		if p*q > 0 {
			c++
		}
		a++
	}

	Print(c)
}
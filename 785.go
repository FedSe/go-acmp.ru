package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	for a <= b {
		p := 1
		t := a
		for t > 0 {
			p *= 10
			t /= 10
		}
		if a*a%p == a {
			Println(a)
		}
		a++
	}
}
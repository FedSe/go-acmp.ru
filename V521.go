package main
import . "fmt"
func main() {
	var p, k, f int
	Scan(&p, &k)

	for p <= k {
		n := p
		for n != 2 {
			f++
			if n%2 < 1 {
				n /= 2
			} else {
				n = n*3 + 1
			}
		}
	p++
	}

	Print(f)
}
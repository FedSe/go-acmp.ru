package main
import . "fmt"
func main() {
	d := 0
	p := 1

	Scan(&d)
	for p <= d {
		p <<= 1
	}

	Print(d ^ (p - 1))
}
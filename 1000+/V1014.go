package main
import . "fmt"
func main() {
	i := 2
	r := 1
	n := 0

	Scan(&n)
	for i < n {
		i++
		r = (r*i + 1 - i&1*2) % (1e9 + 9)
	}

	Print(r)
}
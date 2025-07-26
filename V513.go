package main
import . "fmt"
func main() {
	n := 0
	i := 0
	r := 1
	Scan(&n)

	for i < n {
		r *= 2
	i++
	}

	Print(r - n - 1)
}
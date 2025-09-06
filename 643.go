package main
import . "fmt"
func main() {
	n := 0

	Scan(&n)
	k := n
	for k > 0 {
		n += k & 1
		k /= 2
	}

	Print(n)
}
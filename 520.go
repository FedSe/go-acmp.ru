package main
import . "fmt"
func main() {
	n := 0

	Scan(&n)
	a := n / 144
	b := n % 144 / 12
	n %= 12

	if 1026 < n*105 {
		b++
		n = 0
	}

	if 11401 < b*1025+n*105 {
		a++
		b = 0
		n = 0
	}

	Print(a, b, n)
}
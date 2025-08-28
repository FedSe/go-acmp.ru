package main
import . "fmt"
func main() {
	n := 0
	s := "NO"

	Scan(&n)

	for n % 5 < 1 {
		n /= 5
	}
	for n & 1 < 1 {
		n /= 2
	}
	if n > 1 {
		s = "YES"
	}

	Print(s)
}
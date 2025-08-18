package main
import . "fmt"
func main() {
	n := 0
	s := "No"
	Scan(&n)

	for n%2 < 1 {
		n /= 2
	}

	for n%5 < 1 {
		n /= 5
	}

	if n < 2 {
		s = "Yes"
	}

	Print(s)
}
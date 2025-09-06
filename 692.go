package main
import . "fmt"
func main() {
	n := 0
	s := "NO"

	Scan(&n)
	if n&(n-1) < 1 && n > 0 {
		s = "YES"
	}

	Print(s)
}
package main
import . "fmt"
func main() {
	n := 0
	s := "NO"

	Scan(&n)
	switch n {
	case 3, 4, 6:
		s = "YES"
	}

	Print(s)
}
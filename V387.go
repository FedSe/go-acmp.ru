package main
import . "fmt"
func main() {
	n := 0
	k := 0
	s := ""

	Scan(&n)
	for 0 < n {
		Scan(&s)
		if s[0] == s[3] {
			k++
		}
		n--
	}

	Print(k)
}
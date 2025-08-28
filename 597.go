package main
import . "fmt"
func main() {
	var r, n, m int
	s:="YES"
	Scan(&r, &n, &m)

	if r < n + m {
		s="NO"
	}

	Print(s)
}
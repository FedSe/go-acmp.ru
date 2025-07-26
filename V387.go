package main
import . "fmt"
func main() {
	var n, k, i int
	Scan(&n)
	s:=""
	for i < n {
		Scan(&s)
		if s[0] == s[3] { k++ }
	i++
	}

	Print(k)
	
}
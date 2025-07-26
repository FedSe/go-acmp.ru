package main
import . "fmt"
func main() {
	var m, s, n, a int
	Scan(&n)
	for 0 < n {
		Scan(&a)
		s++
		if a < 1 { s = 0 }
		if s > m { m = s }
	n--
	}
 
	Print(m)
}
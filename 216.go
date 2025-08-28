package main
import . "fmt"
func main() {
	var n, s, m, v int
	Scan(&n)

	for 0 < n {
		Scan(&v)
		s += v
		if v > m {
			m = v
		}
	n--
	}

	n = s-m
	s /= 2
	if s < n {
		n = s
	} 

	Print(n)

}
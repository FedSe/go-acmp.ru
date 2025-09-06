package main
import . "fmt"
func main() {
	var n, m, v, k int
	s := "YES "

	Scan(&n, &v, &k)
	for 0 < n {
		if v > 0 {
			m += v
		}
		v -= k
		n--
	}
	if v+k < 1 {
		s = "NO "
	}

	Print(s, m)
}
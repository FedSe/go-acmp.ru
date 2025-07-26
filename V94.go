package main
import . "fmt"
func main() {
	var n, m, k, c int
	Scan(&n, &m, &k)

	if m > n && k >= n {
		Print("NO")
		return
	}

	for {
		c++
		m -= n
		if m < 1 {
			Print(c)
			break
		}
		m += k
	}
}
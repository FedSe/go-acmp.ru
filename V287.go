package main
import . "fmt"
func main() {
	a := map[string]int{}
	var n, m, i int
	s:=""
	Scan(&n, &m, &s)

	for i <= n - m {
		a[s[i:i+m]] = 1
	i++
	}

	Print(len(a))
}
package main
import . "fmt"
func main() {
	var (
		n, m, i int
		v = "YES"
		s = v
		a = make(map[any]int)
	)

	Scan(&n, &m, &s)
	for i <= n-m {
		a[s[i:i+m]] = 1
	i++
	}

	if len(a) > n - m {
		v = "NO"
	}

	Print(v)
}
package main
import . "fmt"
func main() {
	var (
		n, m, i int
		a       = map[any]int{}
		s       = ""
	)

	Scan(&n, &m, &s)
	for i <= n-m {
		a[s[i:i+m]] = 1
		i++
	}

	Print(len(a))
}
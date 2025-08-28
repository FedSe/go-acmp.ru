package main
import . "fmt"
func main() {
	var (
		u       [128]int
		k, l, t int
		s       = ""
	)

	Scan(&k, &k, &s)
	for r, v := range s {
		u[v]++
		for u[v] > k {
			u[s[l]]--
			l++
		}

		t += r - l + 1
	}

	Print(t)
}
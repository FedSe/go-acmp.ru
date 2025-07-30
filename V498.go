package main
import . "fmt"
func main() {
	var (
		n, k, c int
		u       [9]int
		F       func(int, int)
	)

	Scan(&n, &k)
	F = func(p, l int) {
		if p == n {
			c++
			return
		}
		i := 1
		for i <= n {
			if u[i-1] < 1 {
				w := l - i
				if w < 0 {
					w = -w
				}
				if p < 1 || w <= k {
					u[i-1] = 1
					F(p+1, i)
					u[i-1] = 0
				}
			}
			i++
		}
	}

	F(0, 0)
	Print(c)
}
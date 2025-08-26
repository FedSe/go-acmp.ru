package main
import . "fmt"
var (
	n, k, c int
	u       [9]int
)
func F(p, l int) {
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

func main() {
	Scan(&n, &k)
	F(0, 0)
	Print(c)
}
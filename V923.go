package main
import . "fmt"
var m = map[int]int{3: 1}
func f(n int) int {
	a := 0
	if n > 2 {
		r, o := m[n]
		a = r
		if !o {
			m[n] = f(n/2) + f(n-n/2)
			a = m[n]
		}
	}
	return a
}
func main() {
	n := 0
	Scan(&n)
	Print(f(n))
}
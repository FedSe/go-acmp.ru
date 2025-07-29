package main
import . "fmt"
func main() {
	m := map[int]int{1: 1}
	n := 0
	a := 1
	i := 2

	Scan(&n)
	for i <= n {
		if m[i] > 0 {
			a += 3
		} else {
			a += 2
		}
		m[a] = 1
		i++
	}

	Print(a)
}
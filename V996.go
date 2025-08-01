package main
import . "fmt"
func main() {
	m := map[int]int{1: 1}
	n := 0
	a := 1
	i := 1

	Scan(&n)
	for i < n {
		i++
		a += 2
		if m[i] > 0 {
			a++
		}
		m[a] = 1
	}

	Print(a)
}
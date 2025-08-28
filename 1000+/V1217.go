package main
import . "fmt"
func main() {
	var (
		a             [1e3]int
		n, i, j, x, q int
	)

	Scan(&n, &x)
	a[0] = x
	y := x
	for i < n-1 {
		i++
		Scan(&q)
		if q > x {
			x = q
		}
		if q < y {
			y = q
		}
		a[i] = q
	}

	for j < n {
		q = a[j]
		if q == x {
			q = y
		}
		Println(q)
		j++
	}
}
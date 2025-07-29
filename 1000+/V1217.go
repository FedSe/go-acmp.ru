package main
import . "fmt"
func main() {
	var (
		a       [1e3]int
		n, i, j, x int
	)
	
	Scan(&n, &x)
	a[0] = x
	y := x
	for i < n-1 {
		i++
		Scan(&a[i])
		if a[i] > x {
			x = a[i]
		} else if a[i] < y {
			y = a[i]
		}
	}

	for j < n {
		if a[j] == x {
			a[j] = y
		}
		Println(a[j])
		j++
	}
}
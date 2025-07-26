package main
import . "fmt"
func main() {
	var (
		a [1000]int
		n, i, j int
	)
	Scan(&n, &a[0])
	x := a[0]
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
		Print(a[j], " ")
	j++
	}
}
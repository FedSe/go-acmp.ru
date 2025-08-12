package main
import . "fmt"
func main() {
	var (
		n, c, i int
		x, y    [20]int
	)

	Scan(&n)
	for i < n {
		Scan(&x[i], &y[i])
		t := 0 > 1
		j := 0
		for j < i && !t {
			t = x[j]*y[i] == x[i]*y[j] && x[j] < 0 == (x[i] < 0) && y[j]*y[i] > 0
			j++
		}
		if !t {
			c++
		}
		i++
	}

	Print(c)
}
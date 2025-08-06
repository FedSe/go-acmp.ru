package main
import . "fmt"
func main() {
	var N, x, y int
	Scan(&N, &x, &y)
	if x > y {
		x, y = y, x
	}
	N--

	T := (N * x * y) / (x + y)
	for T/x+T/y < N {
		T++
	}
	Print(T + x)
}
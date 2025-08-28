package main
import . "fmt"
func main() {
	var n, i, j int
	P := Print

	Scan(&n)
	if n > 2 {
		P(494)
		for i < n-3 {
			P(9)
			i++
		}
		P(55)
		for j < n-2 {
			P(0)
			j++
		}
	} else {
		i = 4905
		if n < 2 {
			i = 45
		}
		P(i)
	}
}
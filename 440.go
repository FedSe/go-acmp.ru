package main
import . "fmt"
func main() {
	var (
		X    = 0.
		Y    = X
		p    [5]int
		i, k int
	)

	for i < 5 {
		Scan(&X, &Y)
		x := .0
		j := 0
		for j < 5 {
			a := x - X
			if a*a+Y*Y < 101 {
				p[j] = 1
				break
			}
			x += 25
			j++
		}
		i++
	}

	for i > 0 {
		i--
		if p[i] > 0 {
			k++
		}
	}

	Print(k)
}
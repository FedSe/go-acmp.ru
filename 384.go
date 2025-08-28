package main
import . "fmt"
func main() {
	var a, i, j int
	Scan(&i, &j)

	for i*j > 0 {
		if i < j {
			j %= i
		} else {
			i %= j
		}
	}

	b := 1
	for i+j > 0 {
		a, b = b, (a+b)%1e9
		j--
	}

	Print(a)
}
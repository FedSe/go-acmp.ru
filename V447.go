package main
import . "fmt"
func main() {
	n := 0
	m := 1
	i := 0
	Scan(&n)

	for i < n {
	i++
		m *= i
		for m%10 < 1 {
			m /= 10
		}
		m %= 1e6
	}

	Print(m % 10)
}
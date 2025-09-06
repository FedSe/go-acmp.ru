package main
import . "fmt"
func main() {
	var n, k, i int

	Scan(&n)
	for i < n {
		i++
		r := 0
		l := i
		for l > 0 {
			r = r*10 + l%10
			l /= 10
		}

		if i+r == n {
			k++
		}
	}

	Print(k)
}
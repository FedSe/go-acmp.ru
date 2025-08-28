package main
import . "fmt"
func main() {
	var a, b, d, e, f, i int

	Scan(&a, &b)
	for i < a {
		i++
		j := 0
		for j < b {
			j++
			if i*j%5 < 1 {
				f++
			} else if i*j%3 < 1 {
				e++
			} else if i*j%2 < 1 {
				d++
			}
		}
	}

	Print(`RED :`, d,
		`
GREEN : `, e,
		`
BLUE : `, f,
		`
BLACK : `, a*b-d-e-f)
}
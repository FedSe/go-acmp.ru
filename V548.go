package main
import . "fmt"
func main() {
	a := ""
	b := a
	i := 0
	j := 0
	P := Print

	Scan(&a, &b)
	for i < len(a) && j < len(b) {
		if a[i:]+b[j:] < b[j:]+a[i:] {
			P(a[i] - 48)
			i++
		} else {
			P(b[j] - 48)
			j++
		}
	}

	P(a[i:] + b[j:])
}
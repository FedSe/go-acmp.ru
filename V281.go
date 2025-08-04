package main
import . "fmt"
func main() {
	var n, k, t int

	Scan(&n, &k)
	for k <= n {
		c := 1
		i := 0
		for i < k {
			c = c * (n - i) / (i + 1)
			i++
		}
		t += c
		k++
	}
	Print(t)
}
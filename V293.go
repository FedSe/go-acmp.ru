package main
import . "fmt"
func main() {
	var (
		b, n, j, c, l int
		a             [100]int
		S = Scan
	)

	S(&n)
	for b < n {
		S(&a[b])
		b++
	}

	for j < n {
		S(&b)
		b *= a[j]
		if b > c {
			c = b
			l = j
		}
		j++
	}

	Print(l + 1)
}
package main
import . "fmt"
func main() {
	var (
		b, n, j, c, l int
		a [100]int
	)
	Scan(&n)

	for b < n {
		Scan(&a[b])
	b++
	}

	for j < n {
		Scan(&b)
		b *= a[j]
		if b > c {
			c = b
			l = j
		}
	j++
	}

	Print(l+1)
}
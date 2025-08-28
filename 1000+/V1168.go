package main
import . "fmt"
func main() {
	var (
		a                [2e3]string
		b                [2e3]int
		n, m, x, y, z, i int
	)

	Scan(&n, &m)
	for i < n {
		Scan(&a[i], &x, &y, &z)
		b[i] = x + y + z
		if i > 0 && (b[i] > b[i-1] || b[i] == b[i-1] && a[i] > a[i-1]) {
			c := a[i]
			B := b[i]
			j := i - 1
			for j >= 0 && (B > b[j] || B == b[j] && c > a[j]) {
				a[j+1] = a[j]
				b[j+1] = b[j]
				j--
			}
			a[j+1] = c
			b[j+1] = B
		}
		i++
	}

	m--
	Println(a[m], b[m])
}
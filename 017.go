package main
import . "fmt"
func main() {
	var (
		p, a       [4e4]int
		n, l, i, j int
	)

	Scan(&n)
	for j < n {
		Scan(&a[j])
		j++
	}

	for i < n-1 {
		i++
		for l > 0 && a[l] != a[i] {
			l = p[l]
		}
		if a[l] == a[i] {
			l++
		}
		p[i+1] = l
	}

	for (n-1)%(n-l) > 0 {
		l = p[l]
	}
	Print(n - l)
}
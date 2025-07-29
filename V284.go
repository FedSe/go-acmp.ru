package main
import . "fmt"
func main() {
	var (
		n, i, c int
		a       [1e3]int
	)

	Scan(&n)
	for i < n {
		Scan(&a[i])
		i++
	}

	Scan(&n)
	for 0 < n {
		Scan(&i, &c)
		i--
		for i < c {
			Println(a[i])
			i++
		}
		n--
	}
}
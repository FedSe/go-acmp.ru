package main
import . "fmt"
func main() {
	var (
		n, i, c int
		a       [1e3]int
		S = Scan
	)

	S(&n)
	for i < n {
		S(&a[i])
		i++
	}

	S(&n)
	for 0 < n {
		S(&i, &c)
		i--
		for i < c {
			Println(a[i])
			i++
		}
		n--
	}
}
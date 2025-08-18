package main
import . "fmt"
func main() {
	n := 0
	t := ""
	w := map[any]int{}

	Scan(&n)
	for 0 < n {
		Scan(&t)
		w[int(t[0])] = 1
		n--
	}

	for w[65+n] > 0 {
		n++
	}

	Print(n)
}
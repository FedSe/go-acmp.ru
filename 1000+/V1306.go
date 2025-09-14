package main
import . "fmt"
func main() {
	var (
		d    []rune
		n, i int
	)

	Scan(&n)
	for i < n {
		d = append(d, 65+rune(i))
		i++
	}
	o := make([]rune, n)
	for n > 0 {
		n--
		if n&1 > 0 {
			o[n] = d[0]
			d = d[1:]
		} else {
			i = len(d) - 1
			o[n] = d[i]
			d = d[:i]
		}
	}

	Print(string(o))
}
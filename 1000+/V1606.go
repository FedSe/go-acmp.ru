package main
import . "fmt"
func main() {
	var (
		d       [10]string
		m, n, i int
		S       = Scan
	)
	for n < 10 {
		S(&d[n])
		n++
	}

	S(&m)
	for i < m {
		S(&n)
		s := Sprint(n)
		m := ""
		for _, c := range s {
			w := d[c-48]
			if w > m {
				m = w
			}
		}
		Println(m)
		i++
	}
}
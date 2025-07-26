package main
import . "fmt"
func main() {
	var h, w, n int
	Scan(&w, &h, &n)
	a := w
	b := h
	if w < h {
		a = h
		b = w
	}
	b *= n

	for a < b {
		m := (a + b) / 2
		if n > (m/w)*(m/h) {
			a = m + 1
		} else {
			b = m
		}
	}

	Print(a)
}
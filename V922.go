package main
import . "fmt"
func main() {
	var t, a, b, c, d, s float64
	Scan(&a, &b, &c, &d, &s)

	if s > c && c <= d {
		Print("NO")
		return
	}

	for {
		if c >= s {
			t += s / c * a
			Printf("%.2f", t)
			break
		}
		t += a + b
		s -= c - d
	}

}
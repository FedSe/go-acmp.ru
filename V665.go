package main
import . "fmt"
func main() {
	h := 0
	m := 0
	f := "%02d:%02d"
	Scanf(f, &h, &m)

T:
	m++
	h += m / 60
	m = m % 60
	h %= 24

	if !(h%10 == m/10 && h/10 == m%10) {
		goto T
	}

	Printf(f, h, m)
}
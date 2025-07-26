package main
import . "fmt"
var h, m int

func main() {
	Scanf("%d:%d", &h, &m)

	T:
	m++
	h += m / 60
	m = m % 60
	h %= 24

	if !(h%10 == m/10 && h/10 == m%10) {
		goto T
	}

	Printf("%02d:%02d", h, m)
}
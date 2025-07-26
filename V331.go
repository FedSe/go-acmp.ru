package main
import . "fmt"
func main() {
	var a, b, c, d int

	Scanf("%d:%d", &c, &d)
	Scan(&a, &b)
	a += c
	b += d

	if b > 59 {
		b -= 60
		a++
	}
	
	for a > 23 {
		a -= 24
	}

	Printf("%02d:%02d", a, b)
}
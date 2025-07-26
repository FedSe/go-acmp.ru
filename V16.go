package main
import . "fmt"
var d = 0
func b(p, q int) {
	for
	i := p
	i <= q
	{
		b(i+1, q-i)
	i++
	}
	if q < 1 {
		d++
	}
}
func main() {
	n:=0
	Scan(&n)
	b(1, n)
	Print(d)
}
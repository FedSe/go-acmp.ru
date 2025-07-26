package main
import (
	. "fmt"
	. "math"
	. "os" 
	b "bufio"
)
func main() {
	var n, q, w, i, x, y, l float64
	r := b.NewScanner(Stdin)

	Scanln(&n, &q, &w, &l)

	for r.Scan() {
		Sscan(r.Text(), &x, &y)
		i++
		y = w - y
		x = q - x
		if Sqrt(x*x+y*y) <= l {
			Print(i)
			return
		}
	}

	Print("Yes")
}
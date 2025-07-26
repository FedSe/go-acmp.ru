package main
import . "fmt"
func main(){
	var a, b, h, d, c int
	Scan(&h, &a, &b)

	for {
		d++
		c += a
		if c >= h { break }
		c -= b
	}

	Print(d)
}
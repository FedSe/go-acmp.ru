package main
import . "fmt"
func main(){
	var a, b, c, t int
	Scan(&a, &b, &c, &t)

	d := b * t
	if a < t {
		t -= a
		d = a*b + t*c
	}

	Print(d)
}
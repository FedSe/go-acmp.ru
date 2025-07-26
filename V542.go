package main
import . "fmt"
func main() {
	n:=0
	r:=0
	Scan(&n)

	for n > 0 {
		r *= 2
		r += n % 2
		n /= 2
	}

	Print(r)
}
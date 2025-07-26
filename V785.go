package main
import . "fmt"
func main() {
	a:=0
	b:=0
	Scan(&a, &b)

	for a <= b {
		x := a
		y := a*a
		for x > 0 {
			if x%10 != y%10 {
				goto A
			}
			x /= 10
			y /= 10
		}

		Print(a, " ")
A:
	a++
	}
}
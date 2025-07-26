package main
import ( . "fmt"
. "math"
)

func main() {
	var s, S, d, D, x, y, n, i float64
	Scan(&s, &S, &d, &D, &n)

	for i < n {
	i++
		Scan(&x, &y)
		a := s-x
		b := S-y
		x = d-x
		y = D-y

		if 2*Sqrt(a*a + b*b) <= Sqrt(x*x + y*y) {
			Print(i)
			return
		}
	}

	Print("NO")
}
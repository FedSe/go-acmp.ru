package main
import . "fmt"
func main() {
	var n, m, s int
	P := Println

	Scan(&n, &m)
	for n <= m {
		a := 0
		if n&1 > 0 {
			a = 1
			i := 3
			for i*i <= n {
				if n%i < 1 {
					a = 0
					break
				}
				i += 2
			}
		}
		if n == 2 {
			a = 1
		}
		if a > 0 {
			P(n)
			s++
		}
		n++
	}
	if s < 1 {
		P("Absent")
	}
}
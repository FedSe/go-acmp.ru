package main
import . "fmt"
func F(e int) int {
	if e <= 2 {
		if e == 2 {
			return 1
		}
		return 0
	}
	if e%2 < 1 {
		return 0
	}
	i := 3
	for i*i <= e {
		if e%i < 1 {
			return 0
		}
		i += 2
	}
	return 1
}
func main() {
	var n, m, s int
	P := Println

	Scan(&n, &m)
	for n <= m {
		if F(n) > 0 {
			P(n)
			s += 1
		}
		n += 1
	}
	if s < 1 {
		P("Absent")
	}
}
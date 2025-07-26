package main
import . "fmt"
func F(e int) bool {
	if e <= 2 {
		if e == 2 {
			return true
		}
		return false
	}
	if e%2 < 1 {
		return false
	}
	i := 3
	for i*i <= e {
		if e%i < 1 {
			return false
		}
		i += 2
	}
	return true
}
func main() {
	var n, m, s int
	Scan(&n, &m)
	for n <= m {
		if F(n) {
			Println(n)
			s += 1
		}
		n += 1
	}
	if s < 1 {
		Println("Absent")
	}
}
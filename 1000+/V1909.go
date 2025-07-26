package main
import . "fmt"
func main() {
	var n, s, m, v int
	Scan(&n)
	if n == 1 {
		Print("No")
	} else {
		Println("Yes")
		Scan(&v)
		i := 1
		for i < n {
		i++
			Scan(&m)
			s += i * m
		}
		Print(-s)
		i = 1
		for i < n {
		i++
			Print(" ", i*v)
		}
	}
}
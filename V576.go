package main
import . "fmt"
func main() {
	n:=0
	s:=0
	Scan(&n)
	for 
	i := 1
	i < n
	{
		a := i
		b := n
		for a * b > 1 {
			if a > b { a, b = b, a }
			b %= a
		}
		if a + b < 2 {
			s++
		}
	i++
	}
	Print(s)
}
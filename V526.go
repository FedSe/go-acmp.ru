package main
import . "fmt"
func main() {
	a := 0
	i := 2
	s := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := s
	Scan(&b, &a)
	for i < 37 {
		v := a
		c := ""
		for v > 0 {
			c = string(s[v%i]) + c
			v /= i
		}
		if b == c {
			Print(i)
			return
		}
	i++
	}

	Print(0)
}
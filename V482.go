package main
import . "fmt"
func main() {
	n:=0
	i:=1
	Scan(&n)
	s := "1"
	c := s
	for len(s) <= n {
		s += c
		i++
		c += Sprint(i)
	}
	Print(s[n]-48)
}
package main
import . "fmt"
func main() {
	s:=0
	a:=1
	for a > 0 {
		Scan(&a)
		s += a
	}
	Print(s)
}
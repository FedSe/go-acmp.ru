package main
import . "fmt"
func main() {
	n:=0
	f:=1
	s:=""
	Scan(&n, &s)

	for n > 1 {
		f *= n
		n -= len(s)
	}

	Print(f)
}
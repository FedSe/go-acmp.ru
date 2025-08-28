package main
import . "fmt"
func main(){
	n:=0
	s:=0
	Scan(&n)

	for n > 0 {
		s += n % 2
		n /= 2
	}

	Print(s)
}
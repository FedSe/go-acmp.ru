package main
import . "fmt"
func main(){
	s:=0
	n:=0
	Scan(&n)

	for n > 0 {
		s += n % 10
		n /= 10
	}

	Print(s)
}
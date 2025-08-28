package main
import . "fmt"
func main(){
	var a, m, n int
	l:=30000
	Scan(&n)
	for 0 < n {
		Scan(&a)
		if a > m { m=a }
		if a < l { l=a }
	n--
	}

	Print(l, m)
}
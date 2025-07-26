package main
import . "fmt"
func main(){
	n:=0
	Scan(&n)

	m := n * 6
	k := n / 6
	n %= 6

	if n > 0 {
		k += 7-n
	}

	Print(k, m)
}
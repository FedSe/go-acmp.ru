package main
import . "fmt"
func main() {
	n:=0
	m:=0
	Scan(&n, &m)
	a := n
	for n*m > 0 {
		if n > m {
			n %= m
		} else {
			m %= n
		}
	}
	n+=m
	Print(a / n)

}
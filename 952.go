package main
import . "fmt"
func main(){
	n:=1
	m:=1
	Scan(&n, &m)
	l := m

	if n * m < m {
		Print("Impossible")
	} else {
		if n > m {
			l = n
		}

		if m > 0 {
			n += m - 1
		}

		Print(l, n)
	}
}
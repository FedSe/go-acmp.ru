package main
import . "fmt"
func main() {
	var (
	d [20]int
	n, i int
)
	d[0] = 1
	Scan(&n)
	 
	for j, a := range d[:19] {
		d[j+1] = a * 3
		if n > a { i ++ }
	}

	Print(i)
}
package main
import . "fmt"
func main() {
	m := 0
	x := 1
	k := 1
	Scan(&m)
	for x > 0 {
		Scan(&x)
		if x > m {
			m = x
			k = 1
		} else if x == m {
			k++
		}
	}
	Print(k)
}
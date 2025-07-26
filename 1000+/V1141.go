package main
import . "fmt"
func main() {
	var m, a, b int
	k := 1
	Scan(&b)
	for b > 0 {
		a = b
		Scan(&b)
		if a == b {
			k++
		} else {
			if k > m {
				m = k
			}
		k = 1
		}
	}
	Print(m)
}
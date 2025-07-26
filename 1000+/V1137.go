package main
import . "fmt"
func main() {
	b := 0
	k := 0
	Scan(&b)
	for b > 0 {
		a := b
		Scan(&b)
		if b > a {
			k++
		}
	}
	Print(k)
}
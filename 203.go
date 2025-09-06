package main
import . "fmt"
func main() {
	a := ""
	b := a
	w := 0

	Scan(&a, &b)
	x := len(a)
	if a != b {
		w--
		z := 1
		for z < x {
			c := a[x-z:] + a[:x-z]
			if b == c {
				w = z
				break
			}
			z++
		}
	}

	Print(w)
}
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
		for
		z := 1
		z < x
		{
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
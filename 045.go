package main
import . "fmt"
func main() {
	n := 0
	z := 1
	k := ""

	Scan(&n)
	if n < 2 {
		if n < 1 {
			z = 10
		}
		Print(z)
		return
	}
	z = 9
	for z > 1 {
		for n%z < 1 && n > 1 {
			k = Sprint(z) + k
			n /= z
		}
		z--
	}
	if k == "" || n > 2 {
		k = "-1"
	}

	Print(k)
}
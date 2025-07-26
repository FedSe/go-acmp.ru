package main
import . "fmt"
func F(a int) int {
	b := 1
	z := a
	for z > 1 {
		b *= z
	z--
	}
	return b
}

func main() {
	var n, k, s, i int
	Scan(&n, &k)
	c := F(n) / F(k)
	z := 1
	for i <= n-k {
		s += z*c/F(i)
		z *= -1
	i++
	}

	Print(s)
}
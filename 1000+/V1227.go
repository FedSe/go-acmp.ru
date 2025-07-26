package main
import . "fmt"
func F(n int) int {
	f := 1
	for n > 1 {
		f *= n
		n--
	}
	return f
}

func main() {
	n := 0
	k := 0
	Scan(&n, &k)
	Print(F(n) / F(k) / F(n-k))
}
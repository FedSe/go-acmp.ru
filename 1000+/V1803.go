package main
import . "fmt"
func main() {
	n := 0
	k := 0

	Scan(&n, &k)

	b := 2*n + 1
	f := k / b * 2
	k %= b

	if k < 1 {
		b = n + 1
	} else if k <= n {
		b = k
		f++
	} else {
		b = k - n
		f += 2
	}
    
	Print(f, b)
}
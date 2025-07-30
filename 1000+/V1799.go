package main
import . "fmt"
func main() {
	var n, k, a int

	Scan(&n, &k)
	for n > 0 {
		if n%k > 0 {
			a += n % k
			n -= n % k
		} else {
			n /= k
			a++
		}
	}

	Print(a)
}
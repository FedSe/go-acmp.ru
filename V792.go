package main
import . "fmt"
func main() {
	var n, p, a, b, s, t, k int
	Scan(&n, &p, &a, &b)
	for n > 0 {
		s += n % p
		n /= p
	}
	for a > 0 {
		t += a % b
		a /= b
	}
	if s == t {
		k = s
	}

	Print(k)
}
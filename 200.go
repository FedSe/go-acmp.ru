package main
import . "fmt"
func main() {
	var (
		n, k int
		p    = 1
		a    = 1 << 30
		f    = map[int]int{}
	)

	Scan(&n, &k)
	for p*p < k {
		p++
		for k%p < 1 {
			f[p]++
			k /= p
		}
	}
	if k > 1 {
		f[k]++
	}

	for i, v := range f {
		k = 0
		p = i
		for p <= n {
			k += n / p
			p *= i
		}
		if k/v < a {
			a = k / v
		}
	}

	Print(a)
}
package main
import . "fmt"
func main() {
	var (
		o          [10]int
		d          [1e6 + 1]int
		n, k, l, i int
		f          = 1 << 30
	)

	Scan(&n)
	for l < n {
		Scan(&o[l])
		l++
	}
	Scan(&k)

	for i < k {
		i++
		d[i] = f
	}

	for _, v := range o {
		i = v
		for i <= k {
			z := d[i-v]
			if z < f && z+1 < d[i] {
				d[i] = z + 1
			}
			i++
		}
	}

	k = d[k]
	if k == f {
		k = -1
	}
	Print(k)
}
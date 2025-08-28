package main
import . "fmt"
func main() {
	d := [1e6]int{2: 1}
	i := 2
	k := 0
	p := 0

	Scan(&k, &p)
	for i < k {
		i++
		d[i] = d[i-1]
		if i&1 < 1 && i/2 > 1 {
			d[i] += d[i/2]
			d[i] %= p
		}
	}

	Print(d[k] % p)
}
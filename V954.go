package main
import . "fmt"
func main() {
	var (
		d       [100020][7]int
		k, r, i int
		h       = 10
		M       = int(1e6)
	)
	d[10][0] = 1
	d[10][3] = 1

	Scan(&k)
	for h <= k {
		s := 0
		for s < 6 {
			t := d[h][s] % M
			if t > 0 {
				if s != 2 {
					q := &d[h+1][s+1]
					*q = (*q + t) % M
				}
				if h+10 <= k {
					q := &d[h+10][3-s/9]
					*q = (*q + t) % M
				}
			}
			s++
		}
		h++
	}
	for i < 6 {
		r = (r + d[k][i]) % M
		i++
	}
	Print(r)
}
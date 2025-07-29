package main
import . "fmt"
func main() {
	var (
		w                               [3e4]int
		n, s, a, b, c, d, x, y, r, i, j int
		P                               = Println
	)

	Scan(&n, &a, &c, &b, &d)

	o := b - a
	p := d - c
	k := o*o + p*p

	for i < n {
		i++
		Scan(&x, &y, &r)
		r *= r
		if r*k >= (p*x-o*y+b*c-d*a)*(p*x-o*y+b*c-d*a) {
			w[s] = i
			s++
		}
	}

	P(s)
	for j < s {
		P(w[j])
		j++
	}
}
package main
import . "fmt"
func main() {
	var (
		w [30000]int
		n, s, a, b, c, d, x, y, r, i, j int
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

	Println(s)
	for j < s {
		Print(w[j], " ")
	j++
	}
}
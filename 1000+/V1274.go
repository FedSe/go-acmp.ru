package main
import . "fmt"
func main() {
	var (
		n, s, i int
		w, z, q [1e5]int
	)

	Scan(&n, &s)
	for i < n {
		Scan(&q[i])
		i++
	}

	w[n-1] = q[n-1]
	z[n-1] = q[n-1]
	i = n - 2
	for i >= 0 {
		w[i] = w[i+1]
		if q[i] < w[i] {
			w[i] = q[i]
		}
		z[i] = z[i+1]
		if q[i] > z[i] {
			z[i] = q[i]
		}
		i--
	}

	a := s
	i = s
	for 0 < n {
		n--
		c := s / q[n]
		r := s % q[n]
		o := r + c*w[n]
		if o < a {
			a = o
		}
		x := r + c*z[n]
		if x > i {
			i = x
		}
	}

	Print(a, i)
}
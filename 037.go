package main
import . "fmt"
func main() {
	var (
		n, x, y, z, k int
		q             = .0
		s             = "Yes"
	)

	Scan(&n, &q)
	t := int(q * 1e3)
	for 0 < n {
		Scan(&x, &y, &z, &k)
		if z*z+k*k > t*t*(x*x+y*y)/1e6 {
			s = "No"
		}
		n--
	}

	Print(s)
}
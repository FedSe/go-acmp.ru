package main
import . "fmt"
func main() {
	v := .0
	l := v
	n := 0
    
	Scan(&n, &v, &l)
	l = l / v * 60
	for 0 < n {
		Scan(&v, &v)
		l += v
		n--
	}

	Printf("%.2f", l)
}
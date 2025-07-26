package main
import ."fmt"
func main() {
	v:=.0
	l:=v
	n:=0
	Scan(&n, &v, &l)

	t := l / v * 60

	for 0 < n {
		Scan(&l, &v)
		t += v
	n--
	}

	Printf("%.2f", t)
}
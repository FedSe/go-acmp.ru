package main
import (
	. "fmt"
	b "bufio"
	. "os"
)
func main() {
	a := ""
	u := 1441
	s := b.NewScanner(Stdin)

	Scanf("%s\n", &a)

	for s.Scan() {
		var k, l, x, v int

		t := s.Text()
		f := len(t)-11

		Sscanf(t[f:], "%d:%d %d:%d", &k, &x, &l, &v)
		x += k*60
		v += l*60
		t = t[:f-1]

		k = 1440 - x + v
		if x < v {
			k = v - x
		}

		if k < u {
			a = t
			u = k
		}
	}
	Printf("The fastest train is %s.\nIts speed is %.0f km/h, approximately.", a, 650.0/(float64(u)/60.0))
}
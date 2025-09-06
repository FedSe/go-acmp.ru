package main
import . "fmt"
var d, m, y int

func F(x int) string {
	r := ""
	for x > 0 {
		v := x%d + 55
		if v > 54 && v < 65 {
			v -= 7
		}
		r = string(byte(v)) + r
		x /= d
	}
	return r
}

func main() {
	Scanf("%d/%d/%d", &d, &m, &y)
	d++
	Printf("%s/%s/%s", F(d-1), F(m), F(y))
}
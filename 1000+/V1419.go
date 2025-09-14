package main
import . "fmt"
func main() {
	var a, b, c, d, e int

	Scan(&a, &b, &c, &d, &e)
	for _, v := range []int{e, d, c, b, a} {
		r := 0
		for v > 0 {
			r = r*10 + v%10
			v /= 10
		}
		Println(r)
	}
}
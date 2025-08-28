package main
import . "fmt"
func main() {
	n := 0
	i := 2
	Scan(&n)

	for i < 37 {

		d := 0
		x := n
		for x > 0 {
			t := x % i
			if 1<<t&d > 0 {
				goto A
			}
			d |= 1 << t
			x /= i
		}

		Println(i)
	A:
		i++
	}
}
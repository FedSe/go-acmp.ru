package main
import . "fmt"
func main() {
	n := 0
	i := 0
    
	Scan(&n)
	for n > 0 {
		i++
		x := i
		d := 0
		for x > 0 {
			t := x % 10
			if 1<<t&d > 0 {
				goto A
			}
			d |= 1 << t
			x /= 10
		}
		n--
	A:
	}

	Print(i)
}
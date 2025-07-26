package main
import . "fmt"
func main() {
	var x, y, k, a int
	s := ""

	Scan(&k)
	z := -1
A:
z++
	for z < int(k) {
		Scan(&x, &y, &a)
		f := "0"
		if a == 1 || a == 2 {
			s += "1"
			goto A
		}
		if a > x || a > y {
			s += f
			goto A
		}
		if x%a < 1 {
			if (y-2)%a < 1 {
				f = "1"
			}
			s += f
			goto A
		}
		if x%a == 1 {
			if  y % a + (y-2)%a + (y-1)% a < 1 {
				f = "1"
			}
			s += f
			goto A
		}
		if x%a == 2 {
			if y%a < 1 {
				f = "1"
			}
			s += f
			goto A
		}
		s += "0"
	z++
	}
	Print(s)
}
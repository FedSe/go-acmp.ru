package main
import . "fmt"
func main() {
	n := 0
	k := 1
	s := ""
	for len(s) < 10001 {
		k++
		n = 2
		for n*n <= k {
			if k%n < 1 {
				goto A
			}
		n++
		}
		s += Sprint(k)
A:
	}

	Scan(&n)
	for 0 < n {
		Scan(&k)
		Print(s[k-1]-48)
	n--
	}

}
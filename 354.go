package main
import . "fmt"
func main() {
	n := 0
	i := 1
	s := ""
	Scan(&n)

	for i < n {
	i++
		if i*i > n {
			i = n
		}
		for n % i < 1 {
			Print(s, i)
			s = "*"
			n /= i
		}
	}

}
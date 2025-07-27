package main
import . "fmt"
func main() {
	var (
		n, c, r int
		s       = "YES"
		a       = ""
	)
    
	Scan(&n)
	for n > 0 {
		Scan(&c)
		if c%2 > 0 {
			Println(c)
			r++
		} else {
			a += Sprintln(c)
			r--
		}
		n--
	}

	if 0 < r {
		s = "NO"
	}

	Print(a + s)
}
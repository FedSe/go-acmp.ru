package main
import . "fmt"

func main(){
	var (n, c, r, t int
	a [100] int
	s="\nYES"
)
	Scan(&n)

	for n > 0 {
		Scan(&c)
		if c % 2 > 0 {
			Print(c, " ")
			r++
		} else {
			a[t] = c
			t++
		}
	n--
	}

	Println()
	for n < t {
		Print(a[n], " ")
		n++
	}

	if t < r { s = "\nNO" }

	Print(s)

}
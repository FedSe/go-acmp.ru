package main
import . "fmt"
func main() {
	var (
		a [16]int
		s = ""
		f = map[any]int{}
		i = 0
	)
	
	Scan(&s)
	for _, c := range s {
		f[c]++
	}

	n := len(s)
	a[0] = 1
	for i < n {
		i++
		a[i] = a[i-1] * i
	}

	u := a[n]
	for _, o := range f {
		u /= a[o]
	}

	Print(u)
}
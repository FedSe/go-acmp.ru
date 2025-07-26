package main
import . "fmt"
func main() {
	var (
		a [11] int
		i, b, c int
		s = "Second"
	)

	for b < 11 {
		Scan(&a[b])
	b++
	}

	for i < 11 {
		Scan(&b)
		if a[i] < b { c++ } 
	i++
	}

	if c > 5 { s = "First" }
	Print(s)
}
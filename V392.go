package main
import (
	. "fmt"
	. "os" 
	b "bufio"
	. "strconv"
)
func main() {
	var (
		n, m, i, j int
		a [100000]int
		s = b.NewScanner(Stdin)
	)
	Scan(&n)

	s.Split(b.ScanWords)
	
	for s.Scan() {
		a[i], _ = Atoi(s.Text())
		if a[m] > a[i] {
			m = i
		}
	i++
	}

	for j < n {
		Println(a[(m+j)%n])
	j++
	}
}
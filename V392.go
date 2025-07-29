package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		n, m, i, j int
		a          [1e5]int
		s          = b.NewReader(Stdin)
	)

	Fscan(s, &n)
	for i < n {
		Fscan(s, &a[i])
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
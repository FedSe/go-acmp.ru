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
		S          = Fscan
	)

	S(s, &n)
	for i < n {
		S(s, &a[i])
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
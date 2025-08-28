package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		a    [2e5]int
		f    = map[int]int{}
		n, i int
		s    = b.NewReader(Stdin)
		P    = Println
	)

	Scan(&n)
	for i < n {
		Fscan(s, &a[i])
		i++
	}

	A := a[:n]
	for _, x := range A {
		f[x]++
	}

	i = a[0]
	n = f[i]
	for u, c := range f {
		if c > n || c == n && u < i {
			n = c
			i = u
		}
	}

	for _, x := range A {
		if x != i {
			P(x)
		}
	}

	for _, x := range A {
		if x == i {
			P(x)
		}
	}
}
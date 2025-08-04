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
		S    = Fscan
	)

	S(s, &n)
	for i < n {
		S(s, &a[i])
		i++
	}

	A := a[:n]
	for _, x := range A {
		f[x]++
	}

	i = a[0]
	n = f[i]
	for u, c := range f {
		if c > n || (c == n && u < i) {
			n = c
			i = u
		}
	}

	for _, x := range A {
		if x != i {
			Println(x)
		}
	}

	for _, x := range A {
		if x == i {
			Println(x)
		}
	}
}
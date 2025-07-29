package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
)
func main() {
	var (
		a    []int
		n, i int
		w    = b.NewWriter(Stdout)
		s    = b.NewReader(Stdin)
		S    = Fscan
	)

	S(s, &n)

	for n > 0 {
		S(s, &i)
		a = append(a, i)
		n--
	}
	Ints(a)

	for _, r := range a {
		Fprintln(w, r)
	}

	w.Flush()
}
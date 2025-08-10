package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		n, k, x int
		w       = b.NewReader(Stdin)
		S       = Fscan
	)

	S(w, &n, &k)
	m := map[any]int{}
	for 0 < n {
		S(w, &x)
		m[x] = 1
		n--
	}

	for 0 < k {
		S(w, &x)
		a := "YES "
		if m[x] < 1 {
			a = "NO "
		}
		Print(a)
		k--
	}
}
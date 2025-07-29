package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		v [1e6]int
		s = ""
		i = 1
		w = b.NewWriter(Stdout)
		P = Fprintln
	)

	Scan(&s)
	n := len(s)
	P(w, 0)
	for i < n {
		j := v[i-1]
		for j > 0 && s[i] != s[j] {
			j = v[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		v[i] = j
		P(w, j)
		i++
	}

	w.Flush()
}
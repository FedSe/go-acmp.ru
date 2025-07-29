package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		p [1e5]int
		s = ""
		t = s
		i = 1
	)

	Fscan(b.NewReader(Stdin), &s, &t)
	n := len(t)
	t += "#" + s

	for i < len(t) {
		j := p[i-1]
		for j > 0 && t[i] != t[j] {
			j = p[j-1]
		}
		if t[i] == t[j] {
			j++
		}
		p[i] = j
		i++
	}

	i = 0
	for i < len(s) {
		i++
		if p[n+i] == n {
			Println(i - n)
		}
	}
}
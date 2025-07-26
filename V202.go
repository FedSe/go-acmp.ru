package main
import (
	. "fmt"
	b "bufio"
	. "os"
)

func main() {
	r := b.NewScanner(Stdin)

	r.Scan()
	s := r.Text()
	r.Scan()
	t := r.Text()

	n := len(t)
	t += "#" + s

	var p [100001]int
	i := 1

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

	for
	i = 0
	i < len(s)
	{
	i++
		if p[n+i] == n {
			Print(i-n, " ")
		}
	}
}
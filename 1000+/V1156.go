package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var i, k, l int

	s := b.NewScanner(Stdin)
	s.Buffer(make([]byte, 1<<16), 1<<20)
	s.Scan()

	w := s.Text()
	n := len(w)
	j := 1
	for i < n && j < n && k < n {
		a := w[(i+k)%n]
		b := w[(j+k)%n]
		k++
		if a != b {
			k--
			if a > b {
				i += k + 1
			} else {
				j += k + 1
			}
			if i == j {
				j++
			}
			k = 0
		}
	}
	if i > n {
		i = j
	}

	for l < n {
		Printf("%c", w[(i+l)%n])
		l++
	}
}
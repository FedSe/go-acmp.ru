package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		t    []byte
		s    = b.NewScanner(Stdin)
		g    = -1
		c, i int
	)

	for s.Scan() {
		t = append(append(t, s.Bytes()...), '\n')
	}

	for i < len(t) {
		if g < 0 {
			if t[i] == 47 && t[i+1] == 47 {
				c++
				for i < len(t) && t[i] != '\n' {
					i++
				}
			}
			if t[i] == 123 {
				c++
				i++
				for i < len(t) {
					if t[i] == 125 {
						break
					}
					i++
				}
			}
			if t[i] == 40 && t[i+1] == 42 {
				c++
				i += 2
				for i+1 < len(t) {
					if t[i] == 42 && t[i+1] == 41 {
						i++
						break
					}
					i++
				}
			}
		}
		if t[i] == '\'' {
			g = -g
		}
		i++
	}

	Print(c)
}
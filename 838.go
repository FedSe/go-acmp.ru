package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
func main() {
	r := b.NewScanner(Stdin)
	q := 0
	I := IndexByte

	for r.Scan() {
		s := r.Text()
		i := 0
		for i < len(s) {
			n := int(s[i] - 96)
			if n > 0 && n < 27 {
				q += n%10 + n/10
			}
			n = int(s[i] - 64)
			if n > 0 && n < 27 {
				q += 10 + n%10 + n/10
			}
			if s[i] < 33 {
				q += 4
			}
			n = int(s[i] - 48)
			if n > -1 && n < 10 {
				q += 13 - n
			}
			if s[i] == 46 {
				q += 5
			}
			if s[i] == 59 {
				q += 7
			}
			if s[i] == 44 {
				q += 2
			}
			if I("=+-'\"", s[i]) != -1 {
				q += 3
			}
			if s[i] == 40 || s[i] == 41 {
				q++
			}
			if I("{}[]<>", s[i]) != -1 {
				q += 8
			}
			i++
		}
	}

	Print(q)
}
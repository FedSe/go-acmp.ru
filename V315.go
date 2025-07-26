package main
import (
	. "fmt"
	. "os"
	b "bufio"
)
func main() {
	a := -1
	p := b.NewScanner(Stdin)
	p.Scan()
	s := p.Text()
	m := rune(s[0])

	for _, c := range s {
		if c > m {
			m = c
		}
		if !( (c > 64 && c < 91) || (c > 47 && c < 58) ) {
			goto A
		}
	}

	a = int(m-54)
	if m > 47 && m < 58 {
		a += 7
		if m == 48 {
			a = 2
		}
	}
A:
	Print(a)
}
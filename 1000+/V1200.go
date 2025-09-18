package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
func main() {
	s := b.NewScanner(Stdin)
	k := 0
	Scanln(&k)
	for 0 < k {
		s.Scan()
		e := Fields(s.Text())
		l := len(e)
		for 0 < l {
			l--
			Print(e[l], " ")
		}
		Println()
		k--
	}
}
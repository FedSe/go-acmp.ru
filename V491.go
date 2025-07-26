package main
import (
 . "fmt"
 . "strings"
)
func main() {
	r := "NO SOLUTION"
	s := r
	Scan(&s)
	n := []rune(s)

	for
	i, j := 0, len(n)
	i < j
	{
		j--
		n[i], n[j] = n[j], n[i]
		i++
	}

	if Count(s, string(s[0])) != len(s) {
		r = s
		if s == string(n) {
			r = s[1:]
		}
	}

	Print(r)

}
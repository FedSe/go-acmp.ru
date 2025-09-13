package main
import (
	. "fmt"
	. "strconv"
)
func main() {
	a := 0
	j := 0
	q := map[int]int{}

	Scan(&j)
	s := Sprintf("%b", j)
	for i := range s {
		j = i
		for j < len(s) {
			j++
			m, _ := ParseInt(s[i:j], 2, 0)
			q[int(m)] = 1
		}
	}
	for i := range q {
		a += i
	}

	Print(a)
}
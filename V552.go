package main
import (
	. "fmt"
	. "strconv"
	. "os"
	b "bufio"
)
func main() {
	var (
		c, v [1000]int
		n, i, j int
		s = b.NewScanner(Stdin)
	)

	Scan(&n)
	s.Split(b.ScanWords)

	for s.Scan() {
		v[j], _ = Atoi(s.Text())
	j++
	}

	for i < n {
		j = v[i]
		c[2] += c[1]*j
		c[1] += c[0]*j
		c[0] += j
	i++
	}

	Print(c[2])
}
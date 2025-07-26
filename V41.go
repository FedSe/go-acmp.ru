package main
import (
	. "fmt"
	b "bufio"
	. "os"
	. "sort"
	. "strconv"
)
func main() {
	var (
		a []int
		n = 0
		w = b.NewWriter(Stdout)
		s = b.NewScanner(Stdin)
	)

	s.Split(b.ScanWords)
	Scan(&n)

	for s.Scan() {
		n, _ = Atoi(s.Text())
		a = append(a, n)
	}
	Ints(a)

	for _, r := range a {
		w.WriteString(Sprint(r, " "))
	}
 
	w.Flush()

}
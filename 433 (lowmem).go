package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		O    = 1100000
		u    = make([]int, 2*O)
		p, r int
		s    = b.NewScanner(Stdin)
	)

	s.Buffer(make([]byte, 4096), O)
	u[O] = 1

	s.Scan()
	s.Scan()
	for _, v := range s.Text() {
		p--
		if v < 98 {
			p += 2
		}
		r += u[p+O]
		u[p+O]++
	}

	Print(r)
}
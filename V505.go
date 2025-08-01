package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strconv"
)
func main() {
	s := b.NewScanner(Stdin)
	s.Split(b.ScanWords)

	F := func() int {
		s.Scan()
		v, _ := Atoi(s.Text())
		return v
	}

	var (
		h    = map[any]int{}
		u    []byte
		l, r int
		p    = F()
		o    = F()
		w    = F()
	)

	H := func() {
		u = u[:0]
		i := 0
		for i < p {
			if i > 0 {
				u = append(u, 32)
			}
			u = AppendInt(u, int64(F()), 10)
			i++
		}
	}

	for o > 0 {
		k := F()
		H()
		h[string(u)] = k
		o--
	}

	for w > 0 {
		H()
		v, x := h[string(u)]
		if x {
			l++
			Println(v)
		} else {
			r++
			Println("-")
		}
		w--
	}

	Print("OK=", l, " BAD=", r)
}
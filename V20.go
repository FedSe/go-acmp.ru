package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		n, a, v int
		s = b.NewReader(Stdin)
		m = 1
		x = 1
	)

    Fscan(s, &n)
	if n > 1 {
	    i := n - 2
	    Fscan(s, &a, &v)
		n = v - a
		if n != 0 {
			x++
		}
		if m < x { m = x }
		y := m
		for i > 0 {
		    i--
			a = v
			Fscan(s, &v)
            
			if x == 1 && v-a != 0 || n*(v-a) < 0 {
				x++
				if m < x { m = x }
			} else {
				x = y
			}

			n = v - a
		}
	}

	Print(m)
}
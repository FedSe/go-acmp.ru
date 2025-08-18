package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		a    [3e5]int
		s, i int
		r    = b.NewReader(Stdin)
		w    = b.NewWriter(Stdout)
		F    = func() int {
			x := 0
			p := 1
			h, _ := r.ReadByte()
			for h < 48 || h > 57 {
				if h == 45 {
					p = -p
				}
				h, _ = r.ReadByte()
			}
			for h > 47 && h < 58 {
				x = x*10 + int(h-48)
				h, _ = r.ReadByte()
			}
			return x * p
		}
		n = F()
		m = F()
	)

	for i < n {
		i++
		a[i] = F()
		if a[i] >= m {
			s += a[i]
		}
	}

	n = F()
	for 0 < n {
		i = F()
		x := F()
		q := a[i] >= m
		a[i] += x
		o := a[i] >= m
		if q && !o {
			s -= a[i] - x
		}
		if !q && o {
			s += a[i]
		}
		if q && o {
			s += x
		}
		Fprintln(w, s)
		n--
	}
	w.Flush()
}
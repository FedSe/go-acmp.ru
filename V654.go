package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var n, m, v, t int

	s := b.NewReader(Stdin)
	Fscan(s, &n, &m)

	k := []int{}
	l := m
	for 1 < n {
		Fscan(s, &t)
		if t > l {
			if t > m {
				m = t
			}

			for len(k) > 1 {
				v += k[len(k)-2] - k[len(k)-1]
				l = k[len(k)-2]
				k = k[:len(k)-1]
			}

			v += t - l
			l = m

			if len(k) > 0 {
				k = k[:len(k)-1]
			}
		}

		if t < l {
			k = append(k, t)
			l = t
		}
		n--
	}

	Print(v + m - t)
}
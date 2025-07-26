package main
import (
	b "bufio"
	."fmt"
	. "os"
)
func main() {
	var n, m, v, t int
	s := b.NewScanner(Stdin)
	s.Split(b.ScanWords)

	Scan(&n, &m)

	k := []int{}
	l := m

	i := 1
	for s.Scan() {
		Sscan(s.Text(), &t)
		if t > l {
			if t > m {
				m = t
			}

			for len(k) > 1 {
				l = k[len(k)-1]
				v += k[len(k)-2] - l
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
	i++
	}

	Print(v + m - t)
}
package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
func main() {
	var (
		w                [1e5]int
		r, l, i, P, N, K int
		c                = map[string]int{}
		s                = b.NewReader(Stdin)
	)

	Scanln(&P, &N, &K)
	z := make([]string, P)
	for l < P {
		z[l], _ = s.ReadString('\n')
		l++
	}

	L, _ := s.ReadString('\n')
	for i, v := range Fields(L) {
		Sscan(v, &w[i])
	}

	for i < P {
		L = z[i]
		l = w[i]
		if r >= N {
			break
		}
		if c[L] < K {
			r++
			c[L]++
			Printf(L+` #%d
`, l)
		}
		i++
	}
}
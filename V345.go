package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		n, k, i int
		r       = b.NewReader(Stdin)
		p       = []string{}
		g       = map[string][]string{}
		F       func(string) int
	)

	Scanln(&n)
	for i < n {
		a, _ := r.ReadString('\n')
		p = append(p, a)
		Fscan(r, &k)
		r.ReadString('\n')
		var c []string
		for 0 < k {
			b, _ := r.ReadString('\n')
			c = append(c, b)
			k--
		}
		g[a] = c
		r.ReadString('\n')
		i++
	}

	h := make([]int, n)
	for i, z := range p {
		e := map[any]int{}
		F = func(u string) int {
			if u == z {
				return 1
			}
			if e[u] <= 0 {
				e[u] = 1
				for _, v := range g[u] {
					if F(v) > 0 {
						return 1
					}
				}
			}
			return 0
		}
		for _, x := range g[z] {
			if F(x) > 0 {
				h[i] = 1
				break
			}
		}
	}

	for _, r := range h {
		s := `NO
`
		if r > 0 {
			s = `YES
`
		}
		Print(s)
	}
}
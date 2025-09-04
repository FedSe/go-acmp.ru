package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		q [6]string
		d [101][6]int
		a = ' '
		i = 0
		z = 1
		g = b.NewReader(Stdin)
		k = map[rune]int{
			83: 1,
			87: 2,
			69: 3,
			85: 4,
			68: 5}
	)

	for i < 6 {
		v, _ := g.ReadString('\n')
		q[i] = v[:len(v)-2]
		d[1][i] = 1
		i++
	}

	Fscanf(g, "%c %d", &a, &i)
	for z < i {
		z++
		e := 0
		for e < 6 {
			o := 1
			for _, c := range q[e] {
				o += d[z-1][k[c]]
			}
			d[z][e] = o
			e++
		}
	}

	for l, v := range "NSWEUD" {
		if v == a {
			Print(d[i][l])
			break
		}
	}
}
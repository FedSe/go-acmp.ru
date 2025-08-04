package main
import (
	. "fmt"
	. "math/rand"
)

type P struct{ x, y int }

func main() {
	N := 0
	p := []P{}

	Scan(&N)
	for 0 < N {
		for {
			x := Intn(20001) - 1e4
			y := Intn(20001) - 1e4
			v := 1
			for _, i := range p {
				if i.x == x && i.y == y {
					v = 0
				}
			}
			j := 0
			for j < len(p) && v > 0 {
				k := j + 1
				for k < len(p) {
					a := p[j]
					b := p[k]
					c := P{x, y}
					if (b.x-a.x)*(c.y-a.y) == (b.y-a.y)*(c.x-a.x) {
						v = 0
					}
					k++
				}
				j++
			}
			if v > 0 {
				p = append(p, P{x, y})
				break
			}
		}
		N--
	}

	Print(`YES
`)
	for _, i := range p {
		Println(i.x, i.y)
	}
}
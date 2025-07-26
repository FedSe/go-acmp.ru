package main
import . "fmt"

type P struct {
	a [4] string
}

func (p P) F(i, j int) P {
	r := p
	for _, d := range []int{-1, 0, 1} {
		for _, w := range []int{-1, 0, 1} {
			if d*d+w*w < 2 {
				h := i + d
				g := j + w
				if 0 <= h && h < 4 && 0 <= g && g < 4 {
					r.a[h] = r.a[h][:g] + string(217 - r.a[h][g]) + r.a[h][g+1:]
				}
			}
		}
	}
	return r
}

func main() {
	var s P
	i := 0
	for i < 4 {
		Scan(&s.a[i])
	i++
	}

	u := make(map[P]int)
	q := []P{s}
	u[s] = 0
	for len(q) > 0 {
		p := q[0]
		e := 1
		i = 0
		for i < 4 {
			j := 0
			for j < 4 {
				if p.a[i][j] != p.a[0][0] {
					e = 0
				}
			j++
			}
		i++
		}
		if e > 0 {
			Print(u[p])
			return
		}
		q = q[1:]
		i = 0
		for i < 4 {
			j := 0
			for j < 4 {
				y := p.F(i, j)
				if _, ok := u[y]; !ok {
					u[y] = u[p] + 1
					q = append(q, y)
				}
			j++
			}
		i++
		}
	}

	Print("Impossible")
}
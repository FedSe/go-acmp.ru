package main
import . "fmt"
type P struct{ x, y, z int }
func main() {
	var (
		p             []P
		n, x, y, z, i int
		s             = "No"
		e             = map[P]int{}
	)

	Scan(&n)
	for i < n {
		Scan(&x, &y, &z)
		p = append(p, P{x, y, z})
		e[p[i]] = 1
		i++
	}

	o := p[0]
	for _, v := range p {
		w := 1 > 0
		for _, q := range p {
			w = w && e[P{o.x + v.x - q.x, o.y + v.y - q.y, o.z + v.z - q.z}] > 0
		}
		if w {
			s = "Yes"
			break
		}
	}

	Print(s)
}
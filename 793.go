package main
import . "fmt"
func main() {
	v := 0
	for {
		_, e := Scan(&v)
		if e != nil {
			break
		}
		var (
			a []int
			w = v
			i = 3
			u = v
			d = 0
		)

		for u > 0 {
			d += u % 10
			u /= 10
		}

		for w%2 < 1 {
			a = append(a, 2)
			w /= 2
		}

		for i*i <= w {
			for w%i < 1 {
				a = append(a, i)
				w /= i
			}
			i += 2
		}

		if w > 2 {
			a = append(a, w)
		}

		i = 0
		if a[0] != v {
			v = 0
			for _, o := range a {
				u = 0
				for o > 0 {
					u += o % 10
					o /= 10
				}
				v += u
			}
			if d == v {
				i = 1
			}
		}
		Print(i)
	}
}
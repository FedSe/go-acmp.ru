package main
import . "fmt"
func main() {
	var (
		h = make(map[any]int)
		v = make(map[any][]string)
		i = 0
		j = -1
		a = ""
		w = a
	)

	Scan(&i)
	for 0 < i {
		Scanf(`
%s -> %s`, &a, &w)
		v[a] = append(v[a], w)
		i--
	}
	Scan(&a, &w)

	q := []string{a}
	for len(q) > 0 {
		a = q[0]
		q = q[1:]
		if a == w {
			j = h[a]
			break
		}
		for _, x := range v[a] {
			if h[x] < 1 {
				h[x] = h[a] + 1
				q = append(q, x)
			}
		}
	}

	Print(j)
}
package main
import . "fmt"
func main() {
	var (
		r          [260]int
		a          = ""
		b          = a
		w          = a
		i, j, u, q int
	)

	Scan(&a, &b)
	if b > a {
		a, b = b, a
	}

	if a == b && b == "0" {
		Print(0)
		return
	}

	for q < 2 {
		e := []rune(a)
		j = 0
		k := len(e)
		for j < k {
			k--
			e[j], e[k] = e[k], e[j]
			j++
		}
		a, b = b, string(e)
		q++
	}

	for i, v := range a {
		if v == 49 {
			r[i]++
		}
	}
	for i, v := range b {
		if v == 49 {
			r[i]++
		}
	}

	for u < 1 {
		i = 0
		for i < 255 {
			if r[i] > 1 {
				if i < 2 {
					r[1]++
					r[0] -= 2
					if i > 0 {
						r[2]++
						r[0] += 3
						r[1] -= 3
					}
				} else {
					r[i+1]++
					r[i-2]++
					r[i] -= 2
				}
			}
			if r[i] > 0 && r[i+1] > 0 {
				r[i]--
				r[i+1]--
				r[i+2]++
			}
			i++
		}
		q = 1
		j = 0
		for j < 255 {
			q = q | r[j]
			j++
		}
		if q < 2 {
			u = 1
			j = 0
			for j < 255 {
				if r[j]+r[j+1] > 1 {
					u = 0
				}
				j++
			}
		}
	}

	i = 255
	for r[i] < 1 {
		i--
	}

	for i >= 0 {
		w += Sprint(r[i])
		i--
	}

	Print(w)
}
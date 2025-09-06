package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)

	var (
		q = []rune(s)
		x = 0
		b = q[1] > 44
		d = q[0]
		w = q[2]
		p = q[4]
		a = p - d
	)

	for i, c := range q {
		if c > 119 {
			x = i
		}
	}

	if b {
		a = d - p
	}

	if x < 1 {
		a = p - w
		if b {
			a = w + p - 96
		}
	}

	if x > 3 {
		a = d + w - 96
		if b {
			a = d - w
		}
	}

	Print(a)
}
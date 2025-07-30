package main
import . "fmt"
func main() {
	s := ""
	l := 0
	r := 200

	Scan(&s)
	for _, c := range s {
		o := l
		v := r
		i := 200
		j := i
		t := o

		if c != 82 {
			t++
		}
		if i > t {
			i = t
		}

		t = v + 1
		if c != 76 {
			t++
		}
		if i > t {
			i = t
		}

		t = v
		if c != 76 {
			t++
		}
		if j > t {
			j = t
		}

		t = o + 1
		if c != 76 {
			t++
		}
		if j > t {
			j = t
		}

		l, r = i, j
	}
	l++
	if r > l {
		r = l
	}

	Print(r)
}
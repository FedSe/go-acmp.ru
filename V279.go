package main
import . "fmt"
func main() {
	var (
		t       []byte
		s       = ""
		h, b, i int
	)

	Scan(&s)
	for i < len(s) {
		c := s[i]
		if c == 40 || c == 91 {
			b++
			t = append(t, c)
		} else {
			if b == 0 {
				Print(-1)
				return
			}
			b--
			v := len(t) - 1
			o := t[v]
			t = t[:v]
			if (c == 41 && o != 40) || (c == 93 && o != 91) {
				h++
			}
		}
		i++
	}

	if b != 0 {
		h = -1
	}
	Print(h)
}
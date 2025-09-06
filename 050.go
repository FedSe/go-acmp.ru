package main
import . "fmt"
func main() {
	var (
		v       []string
		c, i, l int
		a       = ""
		t       = a
	)

	Scan(&a, &t)
	m := len(t)
	for l < len(t) {
		v = append(v, t)
		t = t[1:] + string(t[0])
		l++
	}

	for i <= len(a)-m {
		for _, s := range v {
			if s == a[i:i+m] {
				c++
				break
			}
		}
		i++
	}

	Print(c)
}
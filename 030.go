package main
import . "fmt"

var (
	v [99]int
	a = ""
	b = a
	i = 48
)

func F() {
	for _, o := range a {
		v[o]++
	}
}

func main() {
	Scan(&a, &b)
	F()
	for a != b {
		w := []byte(a)
		w[7]++
		z := 7
		l := 0
		for l < 5 {
			if int(w[z]) > 57-l&1*4 {
				w[z] = 48
				w[z-1-l&1]++
			}
			z -= 1 + l&1
			l++
		}
		a = string(w)
		F()
	}

	for i < 58 {
		Println(v[i])
		i++
	}
}
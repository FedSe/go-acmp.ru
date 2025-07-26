package main
import . "fmt"
func main() {
	var (
		v [99]int
		a = ""
		b = a
		i = 48
	)
	Scan(&a, &b)
	G := func() {
		for _, o := range a {
			v[o]++
		}
	}
	G()
	for a != b {
		w := []byte(a)
		w[7]++
		if w[7] > 57 {
			w[7] = 48
			w[6]++
		}
		if w[6] > 53 {
			w[6] = 48
			w[4]++
		}
		if w[4] > 57 {
			w[4] = 48
			w[3]++
		}
		if w[3] > 53 {
			w[3] = 48
			w[1]++
		}
		if w[1] > 57 {
			w[1] = 48
			w[0]++
		}
		a = string(w)
		G()
	}

	for i < 58 {
		Println(v[i])
	i++
	}
}
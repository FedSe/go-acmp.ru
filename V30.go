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

		for l, z := range []int{7, 6, 4, 3, 1} {
			if int(w[z]) > 57-l%2*4 {
				w[z] = 48
				w[z-1-l&1]++
			}
		}

		a = string(w)
		G()
	}

	for i < 58 {
		Println(v[i])
		i++
	}
}
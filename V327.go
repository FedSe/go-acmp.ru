package main
import . "fmt"
func L(n int) int {
	l := n / 1e3
	r := n % 1e3

	s := 0
	for l+r > 0 {
		s += l % 10
		l /= 10
		s -= r % 10
		r /= 10
	}

	return s
}

func main() {
	k := 0
	Scan(&k)

	for 0 < k {
		a := "No"
		s := a
		t := 0
		w := 100000
		Scan(&s)
		for _, i := range s {
			t += w * int(i-48)
			w /= 10
		}
		if L(t-1) * L(t+1) == 0 {
			a = "Yes"
		}
		Println(a)
		k--
	}
}
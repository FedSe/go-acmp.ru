package main
import . "fmt"
func L(n int) int {
	l, r, s := n/1e3, n%1e3, 0
	for l+r > 0 {
		s += l%10 - r%10
		l /= 10
		r /= 10
	}
	return s
}

func main() {
	k := 0

	Scan(&k)
	for k > 0 {
		a := "No"
		s := a
		t := 0
		Scan(&s)
		for _, c := range s {
			t = t*10 + int(c-48)
		}
		if L(t-1)*L(t+1) < 1 {
			a = "Yes"
		}
		Println(a)
		k--
	}
}
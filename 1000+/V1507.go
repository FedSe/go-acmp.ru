package main
import . "fmt"
func main() {
	n := 0
	m := 0
	s := 1

	Scan(&n)
	if n < 2 {
		Print(0)
		return
	}

	for s <= 9*n {
		M := 0
		r := s
		i := 0
		for i < n {
			d := r
			if d > 9 {
				d = 9
			}
			M, r = M*10+d, r-d
			i++
		}
		if r < 1 {
			t := make([]int, n)
			t[0] = 1
			r = s - 1
			l := n - 1
			for l > 0 {
				if r > 8 {
					t[l] = 9
					r -= 9
				} else {
					t[l] = r
					r = 0
				}
				l--
			}
			if r > 0 {
				t[0] += r
			}
			if t[0] < 9 {
				N := 0
				for _, d := range t {
					N = N*10 + d
				}
				if M-N > m {
					m = M - N
				}
			}
		}
		s++
	}

	Print(m)
}
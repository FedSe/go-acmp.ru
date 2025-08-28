package main
import . "fmt"
func main() {
	var (
		A, B int
		f    = ""
		s    = map[int]int{}
	)

	Scanf("%d/%d", &A, &B)

	r := A % B
	A /= B

	for r > 0 {
		i, v := s[r]
		if v {
			f = f[:i] + "(" + f[i:] + ")"
			break
		}
		s[r] = len(f)
		r *= 10
		f += Sprint(r / B)
		r %= B
	}

	if f != "" {
		f = "." + f
	}

	Print(A, f)
}
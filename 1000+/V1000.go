package main
import . "fmt"
type S = string

func F(a int) int {
	x := 48
	if a > 9 {
		x += 39
	}
	return x + a
}

func G(a, k int) S {
	s := ""
	for 0 < k {
		y := a & 255
		s = S(F(y>>4)) + S(F(y&15)) + s
		a >>= 8
		k--
	}
	return s
}

func main() {
	n := 0
	a := 0
	Scan(&n)
	for 0 < n {
		i := 0
		Scan(&a)
		s := "ff" + G(a, 8)
		for i < 8 {
			var (
				b, j, l    int
				p, c, k, v uint
				q          = 255 << 55
			)
			for j < i {
				b |= 255 << p & a
				p += 8
				j++
			}
			b |= (1<<(7-i) ^ 255) << p & a
			p += 7
			for l < i {
				b |= 1 << p
				p--
				l++
			}
			j = b << ((7 - i) * 8)
			for k < 64 {
				if 1<<k&j == 0 {
					c = k
				}
				k++
			}
			k = 63 - c
			if 1<<(c-1)&j > 0 {
				j |= 1 << c
			} else {
				c++
				for c < 64 {
					j ^= 1 << c
					c++
				}
				q = 0
			}
			for v < 7-k {
				j = j>>8 | q
				v++
			}
			if a == j {
				w := G(b, i+1)
				if len(w) < len(s) {
					s = w
				}
			}
			i++
		}
		Println(s)
		n--
	}
}
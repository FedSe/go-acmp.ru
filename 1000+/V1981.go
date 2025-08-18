package main
import . "fmt"
func main() {
	a := 0
	b := 0
	s := ""

	Scan(&s)
	for i := range s {
		Sscan(s[:i], &a)
		Sscan(s[i:], &b)
		if a > 3 && a < 17 && b > 1 && b < 19 {
			w := 9999
			q := 1
			for q < 10 {
				z := 1
				for z < 10 {
					y := 1
					for y < 10 {
						d := 1
						for d < 10 {
							var e, o, c int
							for _, x := range []int{q, z, y, d} {
								o += x
								if x&1 < 1 {
									o -= x
									e += x
									c++
								}
							}
							if c == 2 && e == a && o == b {
								N := q*1e3 + z*100 + y*10 + d
								if N < w {
									w = N
								}
							}
							d++
						}
						y++
					}
					z++
				}
				q++
			}
			if w < 9999 {
				Print(w)
				break
			}
		}
	}
}
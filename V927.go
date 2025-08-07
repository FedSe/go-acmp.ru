package main
import . "fmt"
func main() {
	var (
		p       [2][11][11]int
		C       = ""
		z, l, i int
		M       = 1000000007
	)

	Scan(&C)
	n := len(C)
	v := make([]byte, n)
	for l < n {
		v[l] = C[n-1-l]
		l++
	}

	p[0][10][10] = 1
	for i < n {
		var w [2][11][11]int
		l = 0
		for l < 2 {
			o := 0
			for o < 11 {
				h := 0
				for h < 11 {
					r := p[l][o][h]
					if r != 0 {
						a := 0
						for a < 10 {
							b := 0
							for b < 10 {
								if (i != n-1 || (a != 0 && b != 0)) &&
									(o == 10 || a != o) &&
									(h == 10 || b != h) {
									t := a + b + l
									if t%10 == int(v[i]-48) {
										e := &w[t/10][a][b]
										*e += r
										*e %= M
									}
								}
								b++
							}
							a++
						}
					}
					h++
				}
				o++
			}
			l++
		}
		p = w
		i++
	}

	for _, e := range p[0] {
		for _, l := range e {
			z += l
			z %= M
		}
	}
	Print(z)
}
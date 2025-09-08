package main
import (
	. "fmt"
	. "io"
	o "os"
)
func main() {
	var (
		q       [255]int
		x       []int
		d, _    = ReadAll(o.Stdin)
		h       = string(d)
		y       = 10
		i, l, k int
		m       = len(h)
		A       = Sscan
	)

	for i < m {
		if h[i] == 60 {
			j := i
			if h[i:i+6] == "<font " {
				z := ""
				for j < m {
					j += 5
					if h[j-5:j] == "size=" && (h[j] == 34 || h[j] == 39) {
						g := h[j]
						j++
						k = j
						for h[j] != g {
							j++
						}
						if j > k {
							z = h[k:j]
							break
						}
					}
					j++
				}
				if z > "" {
					x = append(x, y)
					if z[0] == 43 || z[0] == 45 {
						A(z, &k)
						y += k
					} else {
						A(z, &y)
					}
				}
			}
			if i+6 < m && h[i:i+7] == "</font>" {
				k = len(x) - 1
				y = x[k]
				x = x[:k]
				i += 7
				goto A
			}
			for h[i] != 62 {
				i++
			}
			i++
		A:
		} else {
			if h[i] > 32 {
				q[y]++
			}
			i++
		}
	}

	for l < 52 {
		l++
		if q[l] > 0 {
			Println(l, q[l])
		}
	}
}
package main
import . "fmt"
func main() {
	var (
		z [50000]int
		n, j, i, k int
	)

	Scan(&n)
	m := n
	for j < n*100 {
		Scan(&z[j], &z[j+1], &z[j+2], &z[j+3], &z[j+4])
	j += 100
	}

	for i < n*100 {
		if z[i] > z[i+2] {
			z[i], z[i+2] = z[i+2], z[i]
		}
		z[i] -= z[i+4]
		z[i+2] += z[i+4]
	i++
		if z[i] > z[i+2] {
			z[i], z[i+2] = z[i+2], z[i]
		}
		z[i] -= z[i+3]
		z[i+2] += z[i+3]

		z[i+3] = 0
	i += 99
	}

	for i = 0; i < (n-1)*100; i += 100 {
		for j = i + 100; j < n*100; j += 100 {
			f := &z[i+4]
			y := &z[j+4]
			l := z[j]
			if l > z[i] { l = z[i] }
			o := z[j+1]
			if o > z[i+1] { o = z[i+1] }
			g := z[j+2]
			if g < z[i+2] { g = z[i+2] }
			t := z[j+3]
			if t < z[i+3] { t = z[i+3] }
			if z[i+2]-z[i]+z[j+2]-z[j] >= g - l && z[i+3]-z[i+1]+z[j+3]-z[j+1] >= t - o {
				if *f == 0 && *y == 0 {
					m--
					k++
					*f = k
					*y = k
					continue
				}
				if *f == 0 || *y == 0 {
					m--
					if *f < *y { *f = *y }
					if *y < *f { *y = *f }
					continue
				}
				a := *f
				b := *y
				if a != b {
					m--
					c := a
					if c < b { c = b }
					for
					v := 0
					v < n*100
					{
						f = &z[v+4]
						if *f == a || *f == b {
							*f = c
						}
					v += 100
					}
				}
			}
		}
	}

	Print(m)
}
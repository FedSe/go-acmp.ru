package main
import . "fmt"
func main() {
	var (
		a          [8][8]byte
		n, i, z, j int
		s          = ""
	)

	Scan(&n)
	for i < 8 {
		j = i % 2
		for j < 8 {
			a[i][j] = 45
			if i < 3 {
				a[i][j] = 119
			}
			if i > 4 {
				a[i][j] = 98
			}
			j += 2
		}
		i++
	}

	for z < n {
		Scan(&s)
		for len(s) > 4 {
			y := int(s[0] - 97)
			x := int(s[1] - 49)
			b := a[x][y]
			i = int(s[3] - 97)
			j = int(s[4] - 49)
			v := -1
			if i-y > 0 {
				v = 1
			}
			w := -1
			if j-x > 0 {
				w = 1
			}
			o := x
			p := j
			if o > j {
				o = j
				p = x
			}
			for o <= x && x <= p {
				a[x][y] = 45
				y += v
				x += w
			}
			s = s[3:]
			if j == 7 && b == 119 {
				b = 87
			}
			if j < 1 && b == 98 {
				b = 66
			}
			a[j][i] = b
		}
		z++
	}

	i = 8
	for i > 0 {
		i--
		j = 0
		for j < 8 {
			b := a[i][j]
			if b < 1 {
				b = 46
			}
			Printf("%c", b)
			j++
		}
		Println()
	}
}
package main
import . "fmt"
func main() {
	var (
		s          = ""
		x          = s
		t, I, a, b int
		P          = Print
	)

	Scan(&s)
	for i, v := range s {
		if v < 48 {
			x = s[t:i]
			t = i + 1
		}
	}

	for I < 2 {
		a = b
		b = 0
		i := 0
		j := len(x)
		m := map[byte]int{77: 1e3, 68: 500, 67: 100, 76: 50, 88: 10, 86: 5, 73: 1}
		for i < j {
			v := m[x[i]]
			b += v
			if i < j-1 && m[x[i+1]] > v {
				b -= v * 2
			}
			i++
		}
		I++
		x = s[t:]
		if b < 1 {
			P("ERROR")
			return
		}
	}

	t = a
	z := b
	for z > 0 {
		t, z = z, t%z
	}

	b /= t
	z = b
	for I > 0 {
		s = x
		x = ""
		for i, v := range []int{900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
			for b >= v {
				x += []string{"CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}[i]
				b -= v
			}
		}
		b = a / t
		I--
	}

	P(x)
	if z > 1 {
		P("/" + s)
	}
}
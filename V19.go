package main
import . "fmt"
var (
    i, v, k, m, y, r int
    V, G, F, L, K [56]int
    s, d, f string
)
func main() {
	Scan(&s, &d, &f)
	a := int(s[1])-48
	b := int(d[1])-48

	p := int(s[0] - 64)*10
	o := int(d[0] - 64)*10
	h := int(f[0] - 64)*10 + int(f[1])-48

	for i < 8 {
		v += 10
		if v != o {
			L[i] = v+b
		}
		if v != p {
			F[i] = v+a
		}
	i++
	}

	v = 0
	for i < 16 {
		v++
		if v != a {
			F[i] = v+p
		}
		if v != b {
			L[i] = v+o
		}
	i++
	}

	a += p
	o += b

	for _, r := range []int{11,9,-11,-9} {
		v = 0
		for i < 32 {
			v += r
			F[i] = a+v
			if F[i] % 10 % 9 < 1 || F[i] < 11 || F[i] > 88 {
				F[i] = 0
				break
			}
		i++
		}
	}

	for _, r := range []int{h+12,h+21,h+19,h+8,h-12,h-21,h-19,h-8} {
		if r > 10 && r < 89 && r != a && r != o && r % 10 % 9 > 0 {
			K[m] = r
		}
	m++
	}

	i = 0
	for i < 56 {
		if i < 16 {
			if L[i] == h || L[i] == a {
				L[i] = 0
			}
		}

		if i < 32 {
			if F[i] == h || F[i] == o {
				F[i] = 0
			}
			V[i] = F[i]
		}

	i++
	}

	i = 32
	for i < 40 {
		V[i] = K[i-32]
	i++
	}

	for i < 56 {
		V[i] = L[i-40]
	i++
	}


	v = 0
	for k < 56 {

		m = 0
		for
		j := 0
		j < v
		{
			if V[k] == G[j] {
				m++
			}
		j++
		}

		if m == 0 {
			v++
			G[v] = V[k]
		}
	k++
	}

	for y <= v {
		if G[y] > 0 {
			r++
		}
	y++ 
	}

	Print(r)
}
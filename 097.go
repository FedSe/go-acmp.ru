package main
import . "fmt"

var (
	t, X, Y, Z, C          [100]int
	n, i, l, x, y, z, p, k int
	c                      = map[int]int{}
)

func H(x int) int {
	if t[x] != x {
		t[x] = H(t[x])
	}
	return t[x]
}

func main() {
	Scan(&n)
	for l < n {
		Scan(&x, &y, &z, &p, &k)
		if x > z {
			x, z = z, x
		}
		if y > p {
			y, p = p, y
		}
		X[l] = x - k
		Y[l] = y - k
		Z[l] = z + k
		C[l] = p + k

		t[l] = l
		l++
	}

	for i < n {
		l = i
		for l < n {
			if X[i] <= Z[l] && X[l] <= Z[i] && Y[i] <= C[l] && Y[l] <= C[i] {
				t[H(l)] = H(i)
			}
			l++
		}
		i++
	}

	for 0 < n {
		n--
		c[H(n)] = 1
	}

	Print(len(c))
}
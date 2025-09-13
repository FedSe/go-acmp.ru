package main
import . "fmt"
func main() {
	var (
		q, m    [1e4]float64
		z       = 0.
		k, i, j int
		x       = -1.
	)

	Scan(&z, &k)
	for i < k {
		Scan(&q[i])
		i++
	}

	z /= 2
	for _, v := range q {
		if v < z && z < v+1 {
			x = v
		}
	}

	m[j] = x
	if x < 0 {
		j++
		x = -1.
		y := x
		for _, v := range q[:k] {
			if v < z {
				x = v
			}
			if v >= z && y < 0 {
				y = v
			}
		}
		m[j] = x
		j++
		m[j] = y
	}

	for _, V := range q[:k] {
		for _, v := range m[:j+1] {
			if V == v {
				Println(V)
				break
			}
		}
	}
}
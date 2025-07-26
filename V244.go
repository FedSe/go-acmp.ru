package main
import . "fmt"
func main() {
	var (
		u [50000]int
		n, k, z int
		m = -1
	)

	Scan(&n, &k)

	for z < n {
		Scan(&u[z])
	z++
	}

	for z = 0; z < k; z++ {
		r := 0
		o := 0
		p := -1
		w := -1

		for
		x := z
		x < n
		{
			if u[x] < 1 {
				r++
				if r < 2 {
					p = x
				}
			} else {
				o++
				if o < 2 {
					w = x
				}
			}
		x += k
		}

		if r * o < 1 {
			continue
		}

		if r * o > r + o || m >= 0 {
			Print("FAIL")
			return
		}

		m = w
		if m > p { m = p }
		if r * o != 1 {
			m = p
			if o == 1 {
				m = w
			}
		}
	}

	Print("OK ", m+1)
}
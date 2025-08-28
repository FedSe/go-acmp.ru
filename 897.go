package main
import . "fmt"
func main() {
	var (
		a [21]int
		n, x int
	)

	Scan(&n)
	for n > 0 {
		var (
			h = 100
			k, i int
			y = 3
			s = "NO "
			t = 1>0
		)

		Scan(&x)
		for x & 1 < 1 {
			i++
			x /= 2
		}
		if i > 0 {
			k++
			a[k] = i
			if i < h {
				h = i
			}
		}
		for x > 1 && y*y <= x {
			i = 0
			for x%y < 1 {
				i++
				x /= y
			}
			if i > 0 {
				k++
				a[k] = i
				if i < h {
					h = i
				}
			}
			y += 2
		}

		i = 0
		for i < k {
		i++
			t = t && a[i]%h < 1
		}
		if t && x < 2 {
			s = "YES "
		}

		Print(s)
	n--
	}
}
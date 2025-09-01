package main
import . "fmt"
func main() {
	var (
		d          = [17][17]int{{1}}
		a, b, i, l int
		r          = 1
	)

	Scanf("%d:%d", &a, &b)
	if b > a {
		a, b = b, a
	}
	if b > 23 {
		b -= 24
		for i < a-24 {
			j := 0
			for j <= b {
				v := d[i][j]
				if i-j < 1 {
					d[i+1][j] += v
				}
				if j-i < 1 {
					d[i][j+1] += v
				}
				j++
			}
			i++
		}
		for l < 24 {
			l++
			r = r * (48 - l + 1) / l
		}
		r *= d[a-25][b]
	} else {
		a = 24 + b
		for i < b {
			i++
			r = r * (a - i + 1) / i
		}
	}

	Print(r)
}
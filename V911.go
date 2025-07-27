package main
import . "fmt"
func main() {
	n := 0
	var c, u, r, l string
	p := []string{"1 ", "3 ", "9 ", "27 ", "81 ", "243 ", "729 ", "2187 ", "6561 ", "19683 ", "59049 ", "177147 ", "531441 ", "1594323 ", "4782969 ", "14348907 ", "43046721 ", "129140163 ", "387420489 ", "1162261467 "}
	Scan(&c, &n)

	for n > 0 {
		b := n % 3
		d := "+"
		if b < 1 {
			d = "0"
		}
		if b > 1 {
			d = "-"
			n++
		}
		l += d
		n /= 3
	}

	for i, n := range l {
		if n == 45 {
			r += p[i]
		}
		if n == 43 {
			u += p[i]
		}
	}

	if c > "Q" {
		r, u = u, r
	}

	Print("L:", r, `
R:`, u)

}
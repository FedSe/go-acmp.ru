package main
import . "fmt"
func main() {
	a := 0
	q := ' '
	Scanf("%c%d", &q, &a)

	a += int(q * 10)

	for _, r := range []int{-981, 2, 7, 4, 16, 4, 7, 2} {
		a += r
		if a > 10 && a < 89 && a%10%9 > 0 {
			Printf(`%c%d
`, a/10+96, a%10)
		}
	}
}
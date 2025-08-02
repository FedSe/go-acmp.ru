package main
import . "fmt"
func main() {
	f := ""
	t := f
	i := 0
	P := Print

	Scan(&f, &t)

	w := int(f[1] - 49)
	x := int(t[0]) - int(f[0])
	y := int(t[1]-49) - w

	a := x
	if a < 0 {
		a = -a
	}
	b := y
	if b < 0 {
		b = -b
	}

	if a < 2 && b < 2 && a+b > 0 {
		P(`King
`)
		i++
	}
	if x == 0 && w+y > 1 && w+y < 4 {
		P(`Pawn
`)
		i++
	}
	if a*b == 2 {
		P(`Knight
`)
		i++
	}
	if x*y == 0 {
		P(`Rook
`)
		i = 6
	}
	if a == b {
		P(`Bishop
`)
		i = 6
	}
	if i > 5 {
		P(`Queen
`)
	}

	if i < 1 {
		P(`Nobody
`)
	}
}
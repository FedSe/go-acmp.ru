package main
import . "fmt"
type S = string
func main() {
	var (
		b = `black
`
		y = `yellow
`
		r = `red
`
		x, c, z S
	)

	Scan(&x, &c, &z)

	x = map[any]S{
		"bbg": b + b + `GREEN`,
		"bbG": b + y + b,
		"byb": r + b + b,
		"rbb": r + y + b,
		"ryb": b + b + `green`,
		"bYb": b + `YELLOW
` + b,
		"Rbb": `RED
` + b + b}[S([]byte{x[0], c[0], z[0]})]
	if x == "" {
		x = "error"
	}

	Print(x)
}
package main
import . "fmt"
func main() {
	var (
		b = `black
`
		y = `yellow
`
		r = `red
`
		x = b
		c = x
		z = x
	)

	Scan(&x, &c, &z)
	t := map[string]string{
		"bbg": b + b + `GREEN`,
		"bbG": b + y + b,
		"byb": r + b + b,
		"rbb": r + y + b,
		"ryb": b + b + `green`,
		"bYb": b + `YELLOW
` + b,
		"Rbb": `RED
` + b + b,
	}

	x = t[string([]byte{x[0], c[0], z[0]})]
	if x == "" {
		x = "error"
	}

	Print(x)
}
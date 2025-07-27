package main
import . "fmt"
func main() {
	a := ""
	b := a
	c := a
	Scan(&a, &b, &c)

	t := map[string]string{
		"bbg": "black\nblack\nGREEN",
		"bbG": "black\nyellow\nblack",
		"byb": "red\nblack\nblack",
		"rbb": "red\nyellow\nblack",
		"ryb": "black\nblack\ngreen",
		"bYb": "black\nYELLOW\nblack",
		"Rbb": "RED\nblack\nblack",
	}

	a = t[string([]byte{a[0], b[0], c[0]})]
	if a == "" {
		a = "error"
	}

	Print(a)
}
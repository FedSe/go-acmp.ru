package main
import . "fmt"
func main() {
	var h, m, a, b int
	f := `%02d:%02d
`
	Scanf(f, &h, &m)
	Scan(&a, &b)

	b += h*60 + m + a*60

	Printf(f, (b/60)%24, b%60)
}
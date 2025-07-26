package main
import . "fmt"
func main() {
	a := ""
	b := a
	c := a
	Scan(&a, &b, &c)

	if a[0] == 'b' && b[0] == 'b' && c[0] == 'g' {
		Print("black\nblack\nGREEN")
	} else if a[0] == 'b' && b[0] == 'b' && c[0] == 'G' {
		Print("black\nyellow\nblack")
	} else if a[0] == 'b' && b[0] == 'y' && c[0] == 'b' {
		Print("red\nblack\nblack")
	} else if a[0] == 'r' && b[0] == 'b' && c[0] == 'b' {
		Print("red\nyellow\nblack")
	} else if a[0] == 'r' && b[0] == 'y' && c[0] == 'b' {
		Print("black\nblack\ngreen")
	} else if a[0] == 'b' && b[0] == 'Y' && c[0] == 'b' {
		Print("black\nYELLOW\nblack")
	} else if a[0] == 'R' && b[0] == 'b' && c[0] == 'b' {
		Print("RED\nblack\nblack")
	} else if a[0] == 'b' && b[0] == 'b' && c[0] == 'G' {
		Print("black\nblack\nGREEN")
	} else {
		Print("error")
	}
}
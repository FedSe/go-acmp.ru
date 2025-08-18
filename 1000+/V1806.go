package main
import . "fmt"
func main() {
	var m, w, h int
	Scan(&m, &w, &h)
	w--
	h--
	Print((w/m + 2) * (h/m + 2))
}
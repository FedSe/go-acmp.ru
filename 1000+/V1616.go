package main
import . "fmt"
func main() {
	w := 0
	k := 0
	Scan(&w, &w, &k)
	Print(k*w - 1)
}
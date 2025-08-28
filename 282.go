package main
import . "fmt"
func main() {
	var w, h, s int
    
	Scan(&w, &h)
	for w > 0 {
		j := 0
		for j < h {
			j++
			s += w * j
		}
		w--
	}
    
	Print(s)
}
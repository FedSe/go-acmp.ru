package main
import . "fmt"
func main() {
	s := "second"
	a := s
	b := s

	Scan(&a, &b)
	
	if a == b {
		s = "draw"
	}
    
	a = string([]byte{a[0], b[0]})
	if a == "pr" || a == "rs" || a == "sp" {
		s = "first"
	}

	Print(s)
}
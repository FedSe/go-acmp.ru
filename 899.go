package main
import . "fmt"
func main() {
	for {
		var (
			s    []rune
			t    = ""
			a    = 0
			_, e = Scan(&t)
		)

		if e != nil {
			break
		}
        
		for _, c := range t {
			q := len(s)
			z := ' '
			if q > 0 {
				z = c - s[q-1]
			}
			if c == 40 || c == 91 || c == 123 {
				s = append(s, c)
			} else if z < 3 {
				s = s[:q-1]
			} else {
				a = 1
				break
			}
		}
        
		if len(s) > 0 {
			a = 1
		}
        
		Print(a)
	}
}
package main
import . "fmt"
func main() {
	s := ""
	t := 0
	o := 0
	Scan(&s)

	for _, c := range s {
		if c < 58 {
			t = t*10 + int(c-48)
		} else {
			a := t
			if a < 1 { a = 1 }
			for
			i := 0
			i < a
			{
				Printf("%c", c)
				o++
				if o > 39 {
					Println()
					o = 0
				}
			i++
			}
		t = 0
		}
	}
}
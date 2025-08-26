package main
import . "fmt"
var (
	d [8]int
	e = map[any]int{}
	s = ""
)
func F(p string) {
	if len(p) == len(s) {
		if e[p] < 1 {
			Println(p)
		}
		e[p] = 1
		return
	}
	for i := range s {
		if d[i] < 1 {
			d[i] = 1
			F(p + string(s[i]))
			d[i] = 0
		}
	}
}

func main() {
	Scan(&s)
	F("")
}
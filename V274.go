package main
import (
	. "fmt"
	r "reflect"
)
func main() {
	c := ""
	v := c
	n := 0
	Scan(&n)

	for 0 < n {
		Scan(&c, &v)
		t := map[any]int{}
		for _, c := range c {
			t[c] = 1
		}

		e := map[any]int{}
		for _, c := range v {
			e[c] = 1
		}

		a := `NO
`
		if r.DeepEqual(t, e) {
			a = `YES
`
		}
		Print(a)
		n--
	}
}
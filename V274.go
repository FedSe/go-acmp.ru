package main
import (
	. "fmt"
	. "reflect"
)
func main() {
	c:=""
	v:=c
	n:=0
	Scan(&n)

	for 0 < n {
		Scan(&c, &v)

		t := map[rune]bool{}
		for _, c := range c {
			t[c] = 1>0
		}

		e := map[rune]bool{}
		for _, c := range v {
			e[c] = 1>0
		}

		a:="NO"
		if DeepEqual(t, e) {
			a = "YES"
		}
		Println(a)
	n--
	}
}
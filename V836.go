package main
import (
 . "fmt"
 . "sort"
)

func main() {
	n:=0
	a:=0
	var v []int
	Scan(&n)

	for 0 < n {
		Scan(&a)
		if a & 1 < a / 64 & 1 {
			v = append(v, a)
		}
	n--
	}
	Ints(v)

	Println(len(v))
	for _, i := range v {
		Print(i, " ")
	}

}
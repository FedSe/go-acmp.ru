package main
import (. "fmt"
. "sort"
)
func main() {
	a,b:= []int{1,2,3},[]int{1,2,3}
	Scan(&a[0], &a[1], &a[2], &b[0], &b[1], &b[2])
	Ints(a)
	Ints(b)
	s:=0
	i:=0
	for i < 3 {
		s+=a[i]*b[i]
	i++
	}

	Print(s)
}
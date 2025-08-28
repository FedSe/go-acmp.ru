package main
import . "fmt"
func main(){
	k:=0
	a:=0
	s:=""
	Scan(&s)

	for _, r := range s {
		if r < 49 { k++ } else { k = 0 }
		if a < k { a = k }
	}

	Print(a)
}
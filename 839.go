package main
import . "fmt"
func main(){
	a:="YES"
	s:=a
	Scan(&s)

	for _, r := range s {
		if r < 49 { a = "NO" }
	}

	Print(a)
}
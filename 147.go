package main
import . "fmt"
func main(){
	o:=0
	f:=0
	a:=1
	Scan(&o)

	for 0 < o {
		f += a
		a = f - a
	o--
	}

	Print(f)
}
package main
import . "fmt"
func main(){
	a:=1
	b:=1
	Scan(&a, &b)

	for b > 0 {
		a %= b
		a, b = b, a
	}

	Print(a)
}
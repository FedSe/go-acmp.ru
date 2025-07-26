package main
import . "fmt"
func main(){
	a:=0
	b:=0
	Scan(&a,&b)
	n:=a*b

	for a > 0 {
		if a < b { a, b = b, a }
		a %= b
	}

	Print(n/(a+b))
}
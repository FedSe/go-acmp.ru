package main
import . "fmt"
func main(){
	var n, b, a, i int
	Scan(&n)

	for i < n {
		Scan(&b)
		a += b
	i++
	}

	Print(a-n+1)
}
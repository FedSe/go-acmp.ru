package main
import . "fmt"
func main(){
	var n, c, i int
	Scan(&n)

	for i <= n {
	i++
		if n % i < 1 {
			c += i
		}
	}

	Print(c)
}
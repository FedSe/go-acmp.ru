package main
import . "fmt"
func main(){
	n:=0
	i:=1
	Scan(&n)

	r := 540
	for i <= n {
		r += 45
		if i != n && i & 1 > 0 {
			r += 5
		} else if i != n && i & 1 < 1 {
			r += 15
		}
	i++
	}

	Print(r / 60, " ", r % 60)
}
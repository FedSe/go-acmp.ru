package main
import . "fmt"
func main(){
	n:=0
	k:=0
	Scan(&k)

	for k > 0 {
		n++
		k -= n+1
	} 

	Print(n)
}
package main
import . "fmt"
func main(){
	n:=0
	i:=1
	Scan(&n)

	for i*i <= n {
		Print(i*i," ")
	i++
	}

}
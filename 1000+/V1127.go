package main
import . "fmt"
func main(){
	n:=0
	t:=1
	Scan(&n)

	for t <= n {
		Print(t," ")
        	t *= 2;
	}

}
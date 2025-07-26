package main
import . "fmt"
func main(){
	a:=1
	Scan(&a)
	Print("The next number for the number ", a, " is ", a+1, ".\nThe previous number for the number ", a, " is ", a-1, ".")
}
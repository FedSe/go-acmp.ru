package main
import . "fmt"
func main(){
	a:=0
	b:=0
	Scan(&a, &b)

        Print((a * b % 109 + 109) % 109 + 1)

}
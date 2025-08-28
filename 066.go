package main
import . "fmt"
func main(){
	s:="qwertyuiopasdfghjklzxcvbnmq"
	n:=s
	i:=0
	Scan(&n)
	for s[i] != n[0] {i++}
	Print(s[i+1:i+2])
}
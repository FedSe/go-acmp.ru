package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)
	Print([]any{"BLACK", "WHITE"}[(s[0]+s[1])&1])
}
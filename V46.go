package main
import . "fmt"
func main() {
	n := 0
	j := 0
	s := "2.71828182845904523536028750"
	Scan(&n)

	if n > 0 {
		for j <= n {
			Print(s[j:j+1])
		j++
		} 
	}

	j = int(s[j]-48)
	if s[n+2] > 52 { j++ }
	Print(j)
}
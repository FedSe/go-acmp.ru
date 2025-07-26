package main
import . "fmt"
func main() {
	a:="NO"
	s:=a
	t:=s
	Scan(&s, &t)

	j := 0
	i := 0
	for i < len(t) {
		if s[j] == t[i] {
			j++
		}

		if j == len(s) {
			a = "YES"
			t = ""
		}
	i++
	}

	Print(a)
}
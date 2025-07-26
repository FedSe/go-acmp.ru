package main
import ."fmt"
func main() {
	n:=0
	Scan(&n)
	for 0 < n {
		a:="YES"
		s:=a
		Scan(&s)
		
		e := len(s)
		if e % 3 > 0 { a = "NO" }
		e /= 3

		l:=0
		for l < 3 {
			for
			j := e*l
			j < e * (1+l)
			{
				if int(s[j]-48) != l { a = "NO" }
			j++
			}
		l++
		}

		Println(a)
	n--
	}
}
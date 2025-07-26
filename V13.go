package main
import . "fmt"
func main(){
	var n, m, i int
	a:=""
	b:=a

	Scan(&a, &b)

	for i < 4 {
		for j := 0
		j < 4
		{
			if a[i] == b[j] {
				if i == j {
					n++
					m--
				}
				m++
			}
		j++
		}
	i++
	}

	Print(n, m)
}
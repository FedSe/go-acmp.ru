package main
import . "fmt"
type P struct { 
	x, y int
}
func F(a, b P) int {
    d := a.x - b.x
    t := a.y - b.y
    return d*d + t*t
}
func main(){
	var p[3] P
	i:=0
	s:="No"

	for i < 3 {
		Scan(&p[i].x, &p[i].y)
	i++
	}
 
	a := F(p[0], p[1])
	b := F(p[1], p[2])
	c := F(p[0], p[2])

	if a + b == c || a + c == b || b + c == a {
		s ="Yes"
	}

	Print(s)

}
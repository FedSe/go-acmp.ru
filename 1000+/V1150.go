package main
import . "fmt"
func main(){
	t:=' '
	for {
		_, e := Scanf("%1c", &t)
		if e != nil || t < 33 { break }
		if (t != 52 && t != 55) {
			Printf("%c", t)
		}
	}
}
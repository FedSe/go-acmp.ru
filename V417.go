package main
import . "fmt"
func main() {
	a := []string{"Tuesday, ", "Wednesday, ", "Thursday, ", "Friday, ", "Saturday, ", "Sunday, ", "Monday, "}
	k := 0
	d := 1
	m := 1
	y := 0
	Scan(&k)
	Print(a[k%7])
	for k > 0 {
		k--
		d++
		if d > 28 {
			if (d > 30 && (m == 4 || m == 6 || m == 9 || m == 11)) || (d == 30 && m == 2) || (d == 29 && m == 2 && y > 0) || d > 31 {
				d = 1
				m++
				if m == 13 {
					m = 1
					y++
				}
			}
		}
	}

	Printf("%02d.%02d", d, m)
}
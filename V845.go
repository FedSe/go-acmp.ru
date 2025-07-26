package main
import . "fmt"
func main() {
	var (
		x = ""
		y = x
		a = []int{6, 2, 4, 8}
		b = []int{1, 3, 9, 7}
		c = []int{6, 4}
		d = []int{1, 7, 9, 3}
		e = []int{6, 8, 4, 2}
		f = []int{1, 9}
		w = 1
	)

	Scan(&x, &y)

	n := int(x[len(x)-1] - 48)
	y = "00" + y
	t := (int(y[len(y)-2]-48))*10 + int(y[len(y)-1]-48)
	switch n {
	case 0, 1, 5, 6:
		w = n
	case 2:
		w = a[t%4]
	case 3:
		w = b[t%4]
	case 4:
		w = c[t%2]
	case 7:
		w = d[t%4]
	case 8:
		w = e[t%4]
	case 9:
		w = f[t%2]
	}


	Print(w)
}
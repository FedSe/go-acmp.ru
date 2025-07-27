package main
import . "fmt"
func main() {
	var (
		s, x, w [1002]int
		n, z, j int
	)
	Scan(&n)

	for z < n {
		z++
		Scan(&s[z])
	}
    
	x[1] = s[1]
	i := 1
    
	for i < n {
		i++
		z = x[i-1] + s[i]
		x[i] = x[i-2] + s[i]
		if z > x[i] {
			x[i] = z
		}
	}
    
	e := x[n]
	b := n
	w[0] = b
    
	for b > 1 {
		b--
		if e == x[b]+s[b+1] {
			j++
			w[j] = b
			e -= s[b+1]
		} else {
			b--
			if b > 0 {
				j++
				w[j] = b
				e -= s[b+2]
			}
		}
	}

	j++
	w[j] = x[n]
    
	for j > -1 {
		Println(w[j])
		j--
	}
}
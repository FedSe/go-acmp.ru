package main
import . "fmt"
var n, k, p int

func F(i, j int, q []byte) {
	if i == k {
		p--
		if p == 0 {
			Print(string(q[:k]))
		}
		return
	}
	for j+k-i <= n {
		q[i] = byte(97 + j)
		j++
		F(i+1, j, q)
	}
}

func main() {
	Scan(&n, &k, &p)
	F(0, 0, make([]byte, k+1))
}
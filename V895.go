package main
import . "fmt"
func main() {
	var (
		i, j int
		s = "Draw"
		a = ""
	)

	for i < 3 {
		w := ""
		Scan(&w)
		a += w
	i++
	}

	for j < 3 {
		i = 3*j
		M := a[j] + a[j+3] + a[j+6] - 164
		Q := a[i] + a[i+1] + a[i+2] - 164
		N := a[4] - 164
		T := N + a[0] + a[8]
		N += a[2] + a[6]

		if M > 99 || Q > 99 || N > 99 || T > 99 {
			s="Win"
		}
		if M == 73 || Q == 73 || N == 73 || T == 73 {
			s="Lose"
		}
	j++
	}
  
	Print(s)
}
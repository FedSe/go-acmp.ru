package main
import . "fmt"

var (
	a                [101][101]int
	b                [101]int
	k, n, s, l, i, j int
	w                = "No"
	S                = Scan
)

func F(j int) {
	if b[j] > 0 {
		return
	}
	b[j] = 1
	k++
	i := 1
	for i <= n {
		if a[j][i] > 0 {
			F(i)
		}
		i++
	}
}

func main() {
	S(&n, &s)

	for {
		S(&i)
		if i < 1 {
			break
		}
		S(&j)
		a[i][j] = 1
	}

	for l < n {
		l++
		if a[s][l] > 0 {
			F(l)
		}
	}

	if k == n-1 {
		w = "Yes"
	}

	Print(w)
}
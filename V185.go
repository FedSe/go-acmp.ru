package main
import . "fmt"

var (
	a [101][101]int
	b [101]int
	k, n, s, l, i, j int
	w = "No"
)
func F(j int) {
	if b[j] > 0 {
		return
	}
	b[j] = 1
	k++
	for
	i := 1
	i <= n
	i++ {
		if a[j][i] > 0 {
			F(i)
		}
	}
}

func main() {
	Scan(&n, &s)

	for {
		Scan(&i)
		if i < 1 {
			break
		}
		Scan(&j)
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
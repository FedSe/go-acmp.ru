package main
import . "fmt"
func F(s string) int {
	k := 0
	for _, i := range s {
		k += int(i - 48)
	}
	if k < 10 {
		return k
	}
	return F(Sprint(k))
}

func main() {
	t := "NO"
	s := t
	Scan(&s)
	for
	i := 1
	i < len(s)
	i++ {
		if F(s[:i]) == F(s[i:]) {
			t = "YES"
		}
	}

	Print(t)
}
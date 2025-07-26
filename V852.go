package main
import . "fmt"
func main() {
	s := ""
	Scan(&s)

	for
	i := 0
	i < len(s)
	{
		o := s[i:i+1]
		for _, n := range "</>abcdefghijklmnopqrstuvwxyz" {
			s = s[:i] + string(n) + s[i+1:]
			j := 0
			if s[0] == 60 || s[len(s)-1] == 62 {
				var e []string
				for j != len(s) {
					if s[j] != 60 {
						j = 0
						break
					}
					j++
					g := 1
					if s[j] == 47 {
						g = 0
						j++
					}
					if s[j] < 97 {
						j = 0
						break
					}
					t := s[j:j+1]
					j++
					for s[j] > 96 {
						t += s[j:j+1]
						j++
					}
					if s[j] != 62 {
						j = 0
						break
					}
					j++
					if g>0 {
						e = append(e, t)
					} else {
						w := len(e)
						if w < 1 || e[w-1] != t {
							j = 0
							break
						}
						e = e[:w-1]
					}
				}

				if j > len(e)*9 {
					Print(s)
					return
				}
			}
		}
		s = s[:i] + o + s[i+1:]
	i++
	}
}
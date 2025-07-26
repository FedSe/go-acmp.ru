package main
import (
    . "fmt"
    . "strings"
    u "unicode"
)
func main() {
    a := ""
    b := a
    Scan(&a, &b)

    m := a + b
    l := 2
    for l > 0 {
        q := ToLower(a)
        i := 0
        for i <= len(a) {
            w := q[:i] + ToLower(b)
            k := 0
            c := 0
            for c < 3 {
                j := Index(w[k:], q)
                if j < 0 {
                    break
                }
                k += j

                e := []rune(w)
                e[i] = u.ToUpper(e[i])
                e[k] = u.ToUpper(e[k])
                s := string(e)
                
                r:=len(s)
                t:=len(m)
                if r < t || (r == t && s < m) {
                    m = s
                }

                k++
                c++
            }
        i++
        }
    l--
    a, b = b, a
    }

    Print(m)
}
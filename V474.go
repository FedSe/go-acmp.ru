package main
import . "fmt"
func main() {
    n := ""
    b := 0
    
    Scan(&n)
    
    w := len(n) - 1
    r := []rune(n)
    
    i := w
    for i >= 0 {
        if r[i] > 48 {
            r[i]--
            break
        }
        r[i] = 57
    i--
    }

    i = 0
    for i < w && r[i] == 48 {
        i++
    }
    
    n = string(r[i:])
    for len(n) > 1 || n[0] != 48 {
        a := 0
        r = []rune(n)
        for i := range r {
            a = a*10 + int(r[i]-48)
            r[i] = rune(a/3 + 48)
            a %= 3
        }

        i = 0
        for i < len(r)-1 && r[i] < 49 {
            i++
        }
        n = string(r[i:])

        if a > 1 {
            b ^= 1
        }
    }
    
    Print(b)
}
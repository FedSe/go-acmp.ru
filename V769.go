package main
import (
    . "fmt"
    . "sort"
)
func main() {
    var n, t, z, c, i int
    
    Scan(&n)
    s := make([]string, n)
    for i < n {
        Scan(&s[i])
    i++
    }

    Strings(s)
    for z < n {
        i = len(s[t])
        for z < n && len(s[z]) >= i && s[z][:i] == s[t] {
            z++
        }
        t = z 
        c++
    }

    Print(c)
}
package main
import (
    . "fmt"
    . "sort"
)
func main() {
    var (
        s[] string
        n = 0
        b = ""
    )
    
    Scan(&n, &b)
    
    for 0 < n {
        s = append(s, Sprint(n))
    n--
    }
    Strings(s)
    
    Print(SearchStrings(s, b) + 1)
}
package main
import (
    . "fmt"
    . "sort"
)
func main() {
    type A struct {
        a, b  int
    }
    n := 0
	i := 0
    Scan(&n)

    s := make([]A, n)

    for i < n {
        Scan(&s[i].a, &s[i].b)
	i++
    }

    Slice(s, func(i, j int) bool {
        return s[i].a < s[j].a
    })

    e := s[1].b
    f := e

    if n != 2 {
        f += s[2].b
    }
    
    i = 3
    for i < n {
        g := e
        if e > f { g = f }
        g += s[i].b
        e = f
        f = g
    i++
    }

    Print(f)
}
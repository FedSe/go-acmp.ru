package main
import (
    . "fmt"
    . "sort"
)
func F(a, b, c int) int {
    H := func(x, y int) int {
        v, p := 0, 0
        for x + y > 0 {
            q := x%10 + y%10
            x, y = x/10, y/10

            m, i := 1, 0
            for i < p {
                m *= 10
                i++
            }
            v += m * q
            p++
            if q > 9 {
                p++
            }
        }
        return v
    }
    return H(H(a, b), c)
}

func main() {
    var (
        a, b, c int
        s []int
        w = "YES "
        r = map[int]struct{}{}
    )
    Scan(&a, &b, &c)

    n := []struct{ x, y, z int }{
        {a, b, c}, {a, c, b},
        {b, a, c}, {b, c, a},
        {c, a, b}, {c, b, a},
    }

    for _, p := range n {
        r[F(p.x, p.y, p.z)] = struct{}{}
    }

    for v := range r {
        s = append(s, v)
    }

    Slice(s, func(i, j int) bool { return s[i] < s[j] })

    if len(r) < 2 { w = "NO " }
    Print(w)
    for _, v := range s {
        Println(v)
    }
}
package main
import . "fmt"
func main() {
    var (
        s [53][53]int
        n, m, v, y, x int
    )
    
    Scan(&n, &m)
    for y < 53 {
        for
        x = 0
        x < 53
        {
            o := 1
            if y <= n && x <= m && x * y > 0 { o = 0 }
            s[y][x] = o
            x++
        }
        y++
    }

    Scan(&n)
    for n > 0 {
        Scan(&m, &y, &x)
        a := &s[y+1][x+1]
        b := &s[y][x+1]
        c := &s[y+1][x]
        d := &s[y][x]
        if m < 2 && *c + *b + *a < 1 {
            *c = 1
            *b = 1
            *a = 1
            v += 3
        }
        if m == 2 && *c + *d + *a < 1 {
            *c = 1
            *d = 1
            *a = 1
            v += 3
        }
        if m == 3 && *d + *b+ *a < 1 {
            *d = 1
            *b = 1
            *a = 1
            v += 3
        }
        if m > 3 && *d + *b + *c < 1 {
            *d = 1
            *b = 1
            *c = 1
            v += 3
        }
        n--
    }

    Print(v)
}
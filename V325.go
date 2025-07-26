package main
import . "fmt"
var k, z, w int
func F(n, x, y int) {
    n++
    if n < 3 {
        i := x - z
        l := y - w
        if i < 0 { i = -i }
        if l < 0 { l = -l }
        if (i == 1 && l == 2) || (i == 2 && l == 1) {
            k = n
            return
        }

        for _, v := range [][2]int{
            {1, 2}, {1, -2}, {-1, 2}, {-1, -2},
            {2, 1}, {2, -1}, {-2, 1}, {-2, -1},
        } {
            q, g := x+v[0], y+v[1]
            if q > 96 && q < 105 && g > 48 && g < 57 {
                F(n, q, g)
            }
        }
    }
}

func main() {
    a := ""
    b := a
    Scan(&a, &b)

    z, w = int(b[0]), int(b[1])

    F(0, int(a[0]), int(a[1]))

    if k < 1 {
        Print("NO")
    } else {
        Print(k)
    }
}
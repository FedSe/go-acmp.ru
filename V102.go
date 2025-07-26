package main
import ."fmt"

func main(){
    var a, b, c, d, e, f, g, h int
    s := "Out"
    Scan(&a, &b, &c, &d, &e, &f, &g, &h)
    
    F := func (a, b, c, d, e, f int) bool {
        return ((c-a)*(f-b)-(d-b)*(e-a))*((c-a)*(h-b)-(d-b)*(g-a)) >= 0
    }
    
    if  F(a,b,c,d,e,f) && F(c,d,e,f,a,b) && F(e,f,a,b,c,d) {
        s = "In"
    }
    
    Print(s)
}
package main
import . "fmt"
var a [102][102]byte
func F(z, x, n int) {
	A := &a[z-1][x+1]
	B := &a[z-1][x]
	C := &a[z][x+1]
	D := &a[z-1][x-1]
	E := &a[z][x-1]
	G := &a[z+1][x+1]
	H := &a[z+1][x]
	J := &a[z+1][x-1]
	if n < 1 {
		if *B != 42 || *C != 42 {
			if *A == 46 {
				*A = 47
				F(z-1, x+1, 0)
			}
			if *A == 92 {
				*A = 88
				F(z-1, x+1, 0)
			}
		}
		if *A + *C < 85 {
			if *B == 46 {
				*B = 92
				F(z-1, x, 1)
			}
			if *B == 47 {
				*B = 88
				F(z-1, x, 1)
			}
		}
		if *A + *B < 85 {
			if *C == 46 {
				*C = 92
				F(z, x+1, 2)
			}
			if *C == 47 {
				*C = 88
				F(z, x+1, 2)
			}
		}
	}
	if n == 1 {
		if *B != 42 || *E != 42 {
			if *D == 46 {
				*D = 92
				F(z-1, x-1, 1)
			}
			if *D == 47 {
				*D = 88
				F(z-1, x-1, 1)
			}
		}
		if *D + *B < 85 {
			if *E == 46 {
				*E = 47
				F(z, x-1, 3)
			}
			if *E == 92 {
				*E = 88
				F(z, x-1, 3)
			}
		}
		if *D + *E < 85 {
			if *B == 46 {
				*B = 47
				F(z-1, x, 0)
			}
			if *B == 92 {
				*B = 88
				F(z-1, x, 0)
			}
		}
	}
	if n == 2 {
		if *H != 42 || *C != 42 {
			if *G == 46 {
				*G = 92
				F(z+1, x+1, 2)
			}
			if *G == 47 {
				*G = 88
				F(z+1, x+1, 2)
			}
		}
		if *G + *H < 85 {
			if *C == 46 {
				*C = 47
				F(z, x+1, 0)
			}
			if *C == 92 {
				*C = 88
				F(z, x+1, 0)
			}
		}
		if *G + *C < 85 {
			if *H == 46 {
				*H = 47
				F(z+1, x, 3)
			}
			if *H == 92 {
				*H = 88
				F(z+1, x, 3)
			}
		}
	}
	if n > 2 {
		if *H != 42 || *E != 42 {
			if *J == 46 {
				*J = 47
				F(z+1, x-1, 3)
			}
			if *J == 92 {
				*J = 88
				F(z+1, x-1, 3)
			}
		}
		if *J + *E < 85 {
			if *H == 46 {
				*H = 92
				F(z+1, x, 2)
			}
			if *H == 47 {
				*H = 88
				F(z+1, x, 2)
			}
		}
		if *J + *H < 85 {
			if *E == 46 {
				*E = 92
				F(z, x-1, 1)
			}
			if *E == 47 {
				*E = 88
				F(z, x-1, 1)
			}
		}
	}
}

func main() {
	var n, m, c, b, i, j int
	s := ""
	Scan(&n, &m)
	m++
	n++
	for i < 1+n {
		a[i][0] = 42
		a[i][m] = 42
	i++
	}
	for j < 1+m {
		a[0][j] = 42
		a[n][j] = 42
	j++
	}

	for
	i = 1
	i < n
	{
		Scan(&s)
		for
		j = 1
		j < m
		{
			a[i][j] = s[j-1]
			if a[i][j] == 88 {
				c = i
				b = j
			}
		j++
		}
	i++
	}

	F(c, b, 0)
	F(c, b, 1)
	F(c, b, 2)
	F(c, b, 3)

	for
	i = 1
	i < n
	{
		for
		j = 1
		j < m
		{
			Printf("%c", a[i][j])
		j++
		}
		Println()
	i++
	}
}
package main
import . "fmt"
func main() {
	n := 0
	s := `###
`
	v := `# #
`
	Scan(&n)
	Print([]string{
		s + v + v + v + s,
		`  #
 ##
# #
  #
  #`,
		s + `  #
###
#  
` + s,
		s + `  #
###
  #
` + s,
		v + v + s + `  #
  #`,
		s + `#  
###
  #
` + s,
		s + `#  
` + s + v + s,
		s + `  #
  #
  #
  #`,
		s + v + s + v + s,
		s + v + s + `  #
` + s}[n])
}
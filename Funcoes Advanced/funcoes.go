package main

import "fmt"



func CalculosMat(n1 , n2 int) (soma int , sub int ) {
	soma = n1 + n2
	sub = n1 -n2 


	return

}


func main() {
	fmt.Println(CalculosMat(10,40))
}
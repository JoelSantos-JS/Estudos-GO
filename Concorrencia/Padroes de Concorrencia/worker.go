package main


import "fmt"



func main() {

}



func fibo(posi int) int {

	if posi <= 1 {
		return posi
	}

	return fibo(posi - 2 ) + fibo(posi - 1)

}
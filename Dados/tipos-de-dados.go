package main

import "fmt"

func main() {
	const numero int = 1000

	fmt.Println(numero)

	// alias 
	// INT32 = RUNE


	const numero3 float64 = 1122.123
	fmt.Println(numero3)

	// bolleano 

	const jurao bool = true

	fmt.Println(jurao)

	// erro

	const erro error = errors.New("Deu erro")

	fmt.Println(erro)

}
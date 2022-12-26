package main

import "fmt"


func main( ){

	// ARITIMETICOS

	soma := 1 +2 
	subtraçao := 1-2
	multiplicaçao := 2*2
	divisao := 6/2
	resto := 10%2

	fmt.Println(soma,subtraçao,multiplicaçao,divisao,resto)
	// FIM DOS ARITIMETICOS

		// ATRIBUIÇAO

		var num1 string = "Joel"
		num := "Joel 2"

		fmt.Println(num1 ,num)

		// FIM DOS ATIBUIÇAO

		// OPERADORES RELACIONAIS

		fmt.Println(1>2) // false
		fmt.Println(2  >= 2) // true
		fmt.Println( 1 == 2) // false
		fmt.Println(1 < 2) // true
		fmt.Println( 1 <=2) // true
		fmt.Println( 1 != 2) // true

		// OPERADORES LOGICOS

		verdade, falso := true, false

		fmt.Println(verdade && falso)
		fmt.Println(verdade || falso )
		fmt.Println(!!verdade)
		

}
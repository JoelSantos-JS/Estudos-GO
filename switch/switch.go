package main

import "fmt"

func diaDaS(numero int) string {
	

	switch numero {
	case 1: 

	return "Domingo"
	 
	case 2: 
		return "Segunda"
		
	case 3: 
	return "terca"
	
	case 4 : 


	default: 
	return "Numero invalido"

	}

	
}

func main() {

	fmt.Println(diaDaS(2))

}
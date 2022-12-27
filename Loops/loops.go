package main

import( 
	"fmt" 
	

)

func main() {
	// i:= 10

	// Metodo 1
	// for i < 20 {
	// 	i++ 
	// 	fmt.Println("Aumentando 1")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	/// METODO 2

	// for j:= 0; j < 10; j += 5 {
	// 	fmt.Println("aumentdo" , j)

	// 	time.Sleep(time.Second)
	// }
	
	// METODO 3
	nome := [3]string{"Joel", "Jose" , "Santos"}

	for _ , none := range nome {
		fmt.Println(none)
	}
}
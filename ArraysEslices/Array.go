package main


import "fmt"


func main() {
	fmt.Println("Joel")

    ///  tipo de array
	var array =  [5]string{"JJ" ,"ZEI" , "SARA" , "MM" , "JJJHJ"}

	fmt.Println(array[2])

	// tipo 2 de array

	array2 := [...]int{1,2,34,}

	fmt.Println(array2)

	// slice

	slice := []int{1,2,34,5,3,6}

	fmt.Println(slice)

	// append : Adicionar coisas no array .

	slice = append(slice, 17)

	fmt.Println(slice)
	/// ARRAY INTERNOS

	slice3 := make([]int32, 10 , 15)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
 
}
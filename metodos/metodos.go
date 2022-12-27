package main

import "fmt"

type usuarios struct {
	nome string
	idade uint
}

func (u usuarios) MaiordeIdade() string {
	 var nao string = "Nao é maaior"

	if u.idade > 10 {

	} else {
	return	nao
	}


	return  "parabens" 
}

func main() {
		

	user := usuarios{"Joel" , 8}
	joel := user.MaiordeIdade()

	fmt.Println(joel)
}
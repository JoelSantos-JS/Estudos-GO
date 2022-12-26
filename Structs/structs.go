package main


import "fmt"


type usuarios struct {
	nome string
	idade int32
}

func main() {
 fmt.Println("")

 var dados usuarios

	dados.nome = "Joel"
	dados.idade = 12
 fmt.Println(dados)

 u2 := usuarios{"jOAO" , 21}
 fmt.Println(u2)
}
package enderecos


import "strings"

func tipoDeEnd(endereco string) string {

	tiposValido := []string{"Rua" , "Avenida" , "Estradas"}

	primeiraPlavaro := strings.Split(endereco, " ")[0]

}
package main


import (
	"fmt"
	"log"
	"bytes"
	"encoding/json")

type gato struct {
	Nome string `json: "nome"`
	Raca string  `json: "raca"`
	Idade uint   `jso: "idade"`
}

func main() {
 	c := gato{"Ba" ,"irtaP", 3}
	fmt.Println(c)


	gatoJson , erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(gatoJson)
	fmt.Println(bytes.NewBuffer(gatoJson))

}
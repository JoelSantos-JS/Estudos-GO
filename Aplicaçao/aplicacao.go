package main

import (
	"os"
	"log"
	"fmt"
	"linha/app"
)


func main() {
	fmt.Println("jOEL")
	aplica := app.Gerar()
if erro := aplica.Run(os.Args); erro != nil {
	log.Fatal(erro)
}
}
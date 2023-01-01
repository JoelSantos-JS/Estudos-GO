package main

import (
	"fmt"
	//"time"
	
)


func main() {
	canal := make(chan string)

	canal <- "ola mundo"

	 mensagem :=  <-canal
	 fmt.Println(mensagem)

}

